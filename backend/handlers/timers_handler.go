package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dvochoa/1up/db"
	"github.com/dvochoa/1up/models"
	"github.com/gin-gonic/gin"
)

var TimerStore db.TimerStore

// GetTimers returns all timers
func GetTimers(c *gin.Context) {
	timers, err := TimerStore.GetTimers(c.Request.Context())
	if err != nil {
		log.Printf("CollectRows error: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get timers"})
		return
	}

	response := models.GetTimersResponse{Timers: timers}
	c.IndentedJSON(http.StatusOK, response)
}

// GetTimerById returns the timer with matching id, if any
func GetTimerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	timer, err := TimerStore.GetTimerById(c.Request.Context(), id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("timer with id=%d not found", id)})
		return
	}

	c.IndentedJSON(http.StatusOK, timer)
}

// CreateTimer inserts a new timer in the db
func CreateTimer(c *gin.Context) {
	var newTimer models.Timer

	if err := c.BindJSON(&newTimer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := TimerStore.CreateTimer(c.Request.Context(), &newTimer)
	if err != nil {
		log.Printf("Error when creating timer %v: %v", newTimer, err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create timer"})
		return
	}

	c.JSON(http.StatusCreated, newTimer)
}
