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
		AllowMethods:     []string{"GET", "POST", "DELETE", "HEAD"},
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
	engine.Use(PrometheusMiddleware())
	engine.GET("/metrics/", gin.WrapH(promhttp.Handler()))

	router := engine.Group("/api/v1/")
	router.Use(JWTAuth((apiCfg)))
	router.POST("order/", apiCfg.CreateOrder)
	router.GET("order/", apiCfg.GetOrders)
	router.GET("order/:id/", apiCfg.GetOrderDetail)
	router.DELETE("order/:id/", apiCfg.DeleteOrder)
}
