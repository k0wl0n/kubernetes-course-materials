package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/k0wl0n/go-ecommerce-product/api"
	"github.com/k0wl0n/go-ecommerce-product/internal/database"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Check the APP_ENV environment variable to determine the environment
	env := os.Getenv("APP_ENV")
	var envFilePath string

	// Set the .env file path based on the environment
	if env == "production" || env == "staging" {
		envFilePath = "/vault/secrets/.env"
	} else {
		// Default to development environment if APP_ENV is not set
		envFilePath = "/app/.env"
	}

	// Load the environment variables from the appropriate .env file
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file from path: %s, error: %v", envFilePath, err)
	}

	engine := gin.Default()

	// Load env varriables
	port := os.Getenv("PORT")
	mode := os.Getenv("GIN_MODE")
	dbURL := os.Getenv("DB_URL")
	redisURL := os.Getenv("REDIS_URL")

	if port == "" {
		panic("Port is not configured")
	}

	if mode == "" || mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	if dbURL == "" {
		panic("DB URL is not configured")
	}

	// Get DB connection
	dbConn := api.GetDBConn(dbURL)
	defer dbConn.Close()

	// Get Redis connection
	redisClient := api.GetRedisClient(redisURL)
	defer redisClient.Close()

	apiCfg := api.APIConfig{
		DB:      dbConn,
		Queries: database.New(dbConn),
		Cache:   redisClient,
	}

	// Initialize prometheus
	httpRequestsTotal := api.GetPromRequestTotal()
	httpRequestDuration := api.GetPromRequestDuration()

	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)

	// Load API routes
	api.GetRoutes(engine, &apiCfg)

	log.Infoln("Product services started!")
	engine.Run(":" + port)
}
