package main

import (
	"log"

	"github.com/dvochoa/1up/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Specify routes
	router.GET("/timers", handlers.GetTimers)
	router.GET("/timers/:id", handlers.GetTimerById)
	router.POST("/timers", handlers.AddTimer)

	// Start the server
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Println("Failed to start the server:", err)
		return
	}
}
