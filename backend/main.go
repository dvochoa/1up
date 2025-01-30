package main

import (
	"fmt"
	"log"

	"github.com/dvochoa/1up/config" // TODO: Why are these path prefixed with github?
	"github.com/dvochoa/1up/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	router := gin.Default()

	// CORS Configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.AllowedOrigins
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true

	// Apply CORS middleware
	router.Use(cors.New(corsConfig))

	// Specify routes
	router.GET("/timers", handlers.TimersHandler)

	// Start the server
	// TODO: What does :8080 mean? Its what I had before
	fullListenAddress := fmt.Sprintf("%s:%s", cfg.ListenAddress, cfg.Port)
	if err := router.Run(fullListenAddress); err != nil {
		log.Fatal("Failed to start the server:", err)
		return
	}
}
