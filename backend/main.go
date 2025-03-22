package main

import (
	"context"
	"log"

	"github.com/dvochoa/1up/db"
	"github.com/dvochoa/1up/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up router
	router := GetRouter()

	// Set up database connection
	connStr := db.GetDatabaseConnection()
	ctx := context.Background()
	timerStore, err := db.NewTimerStore(ctx, connStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer timerStore.CloseConnection(ctx)
	handlers.TimerStore = *timerStore

	// Start the server
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Println("Failed to start the server:", err)
		return
	}
}

func GetRouter() *gin.Engine {
	router := gin.Default()

	// Going to need the id of the requesting user for:
	// 1. Know how to set owner_id in CreateTimer
	// 2. Checking permissions for all

	// Need to get it passed via some request context, should save that for a follow up change and use some
	// hard-coded fake data for now.
	router.GET("/users/:id/timers", handlers.GetTimers)
	router.GET("/timers/:id", handlers.GetTimerDetails)
	router.POST("/timers", handlers.CreateTimer)
	router.PUT("/timers/:id", handlers.UpdateTimerSettings)
	router.DELETE("/timers/:id", handlers.DeleteTimer)

	return router
}
