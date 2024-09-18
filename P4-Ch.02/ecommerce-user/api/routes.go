package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func GetRoutes(engine *gin.Engine, apiCfg *APIConfig) {
	// CORS Config
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1", "http://localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           5 * time.Hour,
	}))

	// Default route for health check
	engine.GET("/health-check/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Up and Running!",
		})
	})

	// Add Rate limiter middleware
	engine.Use(RateLimiter(apiCfg))

	// Add prometheus middleware and route
	engine.Use(PrometheusMonitoring())
	engine.GET("/metrics/", gin.WrapH(promhttp.Handler()))

	// Public API Routes
	pubRouter := engine.Group("/api/v1/")
	authPubRouter := pubRouter.Group("")
	authPubRouter.Use(JWTAuth((apiCfg)))

	// Non auth routes
	pubRouter.POST("register/", apiCfg.Singup)
	pubRouter.POST("login/", apiCfg.Login)
	pubRouter.POST("refresh-token/", apiCfg.RefreshAccessToken)

	// Auth routes
	authPubRouter.GET("profile/", apiCfg.GetUserProfile)
	authPubRouter.PATCH("profile/", apiCfg.UpdateUserProfile)
	authPubRouter.DELETE("profile/", apiCfg.DeleteUserProfile)
	authPubRouter.PUT("change-password/", apiCfg.ChangePassword)

	// Internal API Routes
	pvtRouter := engine.Group("/internal/v1/")
	pvtRouter.Use(InternalAPIAuth(apiCfg))

	pvtRouter.POST("token/", apiCfg.GetUserFromToken)
}
