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

	router.GET("/timers", handlers.GetTimers)
	router.GET("/timers/:id", handlers.GetTimerById)
	router.POST("/timers", handlers.CreateTimer)
	router.PUT("/timers/:id", handlers.UpdateTimer)
	router.DELETE("/timers/:id", handlers.DeleteTimer)

	return router
}
