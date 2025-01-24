package main

import (
	"log"

	"github.com/dvochoa/1up/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS Configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true

	// Apply CORS middleware
	router.Use(cors.New(config))

	// Specify routes
	router.GET("/timers", handlers.TimersHandler)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Println("Failed to start the server:", err)
		return
	}
}
