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

// GetTimerDetails returns extensive details of the timer with matching id, if any
func GetTimerDetails(c *gin.Context) {
	timerId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timerId"})
		return
	}

	// TODO: Do these both in parallel
	timer, err := TimerStore.GetTimer(c.Request.Context(), timerId)
	if err != nil {
		log.Printf("Error when calling TimerStore.GetTimer: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Failed to get info for timer with id=%d", timerId)})
		return
	}

	timerSessions, err := TimerStore.GetTimerSessions(c.Request.Context(), timerId)
	if err != nil {
		log.Printf("Error when calling TimerStore.GetTimerSessions: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Failed to get history for timer with id=%d", timerId)})
		return
	}

	response := models.GetTimerDetailsResponse{
		Timer:         timer,
		TimerSessions: timerSessions,
	}
	c.JSON(http.StatusOK, response)
}

// AddTimerSession writes a new row into the timerSessions table
func AddTimerSession(c *gin.Context) {
	timerSettingId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timerId"})
		return
	}

	var createTimerSessionRequest models.CreateTimerSessionRequest
	if err := c.BindJSON(&createTimerSessionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid createTimerSessionRequest"})
		return
	}

	timerSession, err := TimerStore.AddTimerSession(c.Request.Context(), timerSettingId, createTimerSessionRequest)
	if err != nil {
		log.Printf("Error when committing session for timer %v: %v", timerSettingId, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit session for timer"})
		return
	}

	c.JSON(http.StatusCreated, timerSession)
}

// CreateTimer inserts a new timer in the db
func CreateTimer(c *gin.Context) {
	var createTimerRequest models.CreateTimerRequest
	if err := c.BindJSON(&createTimerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid createTimerRequest"})
		return
	}

	timer, err := TimerStore.CreateTimerSetting(c.Request.Context(), &createTimerRequest)
	if err != nil {
		log.Printf("Error when creating timer %v: %v", createTimerRequest, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create timer"})
		return
	}

	c.JSON(http.StatusCreated, timer)
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
