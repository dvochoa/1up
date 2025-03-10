package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dvochoa/1up/db"
	"github.com/dvochoa/1up/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var TimerStore db.TimerStore

// GetTimers returns all timers
func GetTimers(c *gin.Context) {
	timers, err := TimerStore.GetTimers(c.Request.Context())
	if err != nil {
		log.Printf("Error when calling TimerStore.GetTimers: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get timers"})
		return
	}

	response := models.GetTimersResponse{Timers: timers}
	c.JSON(http.StatusOK, response)
}

// GetTimerById returns the timer with matching id, if any
func GetTimerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	timer, err := TimerStore.GetTimerById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Timer with id=%d not found", id)})
		return
	}

	c.JSON(http.StatusOK, timer)
}

// CreateTimer inserts a new timer in the db
func CreateTimer(c *gin.Context) {
	var newTimer models.Timer

	if err := c.BindJSON(&newTimer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := TimerStore.CreateTimer(c.Request.Context(), &newTimer)
	if err != nil {
		log.Printf("Error when creating timer %v: %v", newTimer, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create timer"})
		return
	}

	c.JSON(http.StatusCreated, newTimer)
}

func DeleteTimer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := TimerStore.DeleteTimer(c.Request.Context(), id)

	if err != nil {
		log.Printf("Error when deleting timer: %v\n", err)
		msg := gin.H{"message": fmt.Sprintf("Failed to delete timer with id=%d", id)}

		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, msg)
		} else {
			c.JSON(http.StatusInternalServerError, msg)
		}
		return
	}

	c.Status(http.StatusNoContent)
}
