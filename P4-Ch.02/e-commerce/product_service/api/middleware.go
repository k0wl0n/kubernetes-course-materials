package api

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-product/shared"
)

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

		// Fetch userID and set to the context
		userID, err := shared.GetUserFromToken(apiCfg.Cache, ctx, authToken[1])
		if err != nil {
			log.Errorln(err)
			ctx.JSON(http.StatusForbidden, gin.H{"message": "Something went"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", userID)
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
func PrometheusMiddleware() gin.HandlerFunc {
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
