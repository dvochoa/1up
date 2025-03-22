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
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId"})
		return
	}

	timers, err := TimerStore.GetTimers(c.Request.Context(), userId)
	if err != nil {
		log.Printf("Error when calling TimerStore.GetTimers: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get timers"})
		return
	}

	response := models.GetTimersResponse{Timers: timers}
	c.JSON(http.StatusOK, response)
}

// GetTimeHistory returns the history of the timer with matching id, if any
func GetTimerHistory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timerId"})
		return
	}

	timerSessions, err := TimerStore.GetTimerProgress(c.Request.Context(), id)
	if err != nil {
		log.Printf("Error when calling TimerStore.GetTimerProgress: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Failed to get history for timer with id=%d", id)})
		return
	}

	response := models.GetTimerHistoryResponse{TimerSessions: timerSessions}
	c.JSON(http.StatusOK, response)
}

// CreateTimer inserts a new timer in the db
func CreateTimer(c *gin.Context) {
	var newTimer models.Timer

	if err := c.BindJSON(&newTimer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := TimerStore.CreateTimerSetting(c.Request.Context(), &newTimer)
	if err != nil {
		log.Printf("Error when creating timer %v: %v", newTimer, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create timer"})
		return
	}

	c.JSON(http.StatusCreated, newTimer)
}

// UpdateTimer updates the data associated with the timer keyed by id
func UpdateTimerSettings(c *gin.Context) {
	var updatedTimer models.Timer

	if err := c.BindJSON(&updatedTimer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := TimerStore.UpdateTimerSettings(c.Request.Context(), &updatedTimer)
	if err != nil {
		log.Printf("Error when updating timer %v: %v", updatedTimer, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to update timer with id=%d", updatedTimer.Id)})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteTimer deletes all data associated with the timer keyed by id
func DeleteTimer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := TimerStore.DeleteTimerSettings(c.Request.Context(), id)

	if err != nil {
		log.Printf("Error when deleting timer with id=%d: %v", id, err)
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
