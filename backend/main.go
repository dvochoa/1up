package main

import (
	"context"
	"log"

	"github.com/dvochoa/1up/db"
	"github.com/dvochoa/1up/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	connStr := db.GetDatabaseConnection()
	ctx := context.Background()
	timerStore, err := db.NewTimerStore(ctx, connStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer timerStore.CloseConnection(ctx)
	handlers.TimerStore = *timerStore

	// Specify routes
	router.GET("/timers", handlers.GetTimers)
	router.GET("/timers/:id", handlers.GetTimerById)
	router.POST("/timers", handlers.CreateTimer)

	// Start the server
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Println("Failed to start the server:", err)
		return
	}
}
