package api

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-user/utils"
)

// Validate the request by checking wheather or not they have the valid JWT access token or not
//
// Token format: Bearer <TOKEN>
func JWTAuth(apiCfg *APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuthToken := ctx.GetHeader("Authorization")

		if headerAuthToken == "" {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "Authentication required"})
			ctx.Abort()
			return
		}

		// Split the token string
		authToken := strings.Split(headerAuthToken, " ")

		// Validate the token string
		if len(authToken) != 2 || authToken[0] != "Bearer" {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "Invalid authentication format"})
			ctx.Abort()
			return
		}

		// Verify the token and get the encoded payload which is the userID string
		claims, err := utils.VerifyToken(authToken[1])
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "Invalid authentication token"})
			ctx.Abort()
			return
		}

		// Check the validity of the token
		if !time.Unix(claims.ExpiresAt.Unix(), 0).After(time.Now()) {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "Authentication token is expired"})
			ctx.Abort()
			return
		}

		// Convert the userID string to UUID
		userID, err := uuid.Parse(claims.Data)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "Invalid authentication token"})
			ctx.Abort()
			return
		}

		// Fetch the user by the ID from cache
		cachedUser, err := RetriveUserFromCache(apiCfg.Cache, ctx, userID.String())
		if err == nil {
			ctx.Set("user", cachedUser)
		} else {
			log.Errorln("Error caught while retrieving user from cache: ", err)

			// Fetch user by ID from DB
			dbUser, err := GetUserByIDFromDB(apiCfg, ctx, userID)
			if err != nil {
				ctx.JSON(http.StatusForbidden, gin.H{"message": "Something went wrong"})
				ctx.Abort()
				return
			}

			user := DatabaseUserToUser(dbUser)

			// Call goroutine to save user details into cache
			go StoreUserToCache(apiCfg.Cache, ctx, user)

			ctx.Set("user", user)
		}

		// Further call the given handler and send the user instance as well
		ctx.Next()
	}
}

// Middleware for checking if a valid API secret is passed or not
func InternalAPIAuth(apiCfg *APIConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiSecret := ctx.GetHeader("Secret")
		internalAPISecret := os.Getenv("INTERNAL_API_SECRET")

		if internalAPISecret == "" {
			panic("Internal API secret is not configured")
		}

		if apiSecret == "" || apiSecret != internalAPISecret {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "Invalid API Secret"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// Prometheus middleware to record HTTP request timings
func PrometheusMonitoring() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Continue processing the request
		c.Next()

		// Collect metrics after the request is processed
		duration := time.Since(startTime).Seconds()

		httpRequestsTotal := GetPromRequestTotal()
		httpRequestDuration := GetPromRequestDuration()

		httpRequestsTotal.WithLabelValues(c.FullPath(), c.Request.Method).Inc()
		httpRequestDuration.WithLabelValues(c.FullPath(), c.Request.Method).Observe(duration)
	}
}

// Middleware for an IP based rate limiting
func RateLimiter(apiCfg *APIConfig) gin.HandlerFunc {
	limiter := redis_rate.NewLimiter(apiCfg.Cache)
	return func(ctx *gin.Context) {
		// Key is based on the client's IP address
		key := ctx.ClientIP()

		// Allow only 10 requests per minute per IP address
		result, err := limiter.Allow(ctx, key, redis_rate.PerMinute(10))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			ctx.Abort()
			return
		}

		if result.Allowed == 0 {
			ctx.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate Limit Exceeded"})
			ctx.Abort()
			return
		}

		// Continue processing the request
		ctx.Next()
	}
}
