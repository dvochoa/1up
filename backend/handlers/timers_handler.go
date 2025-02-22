package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Timer struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type GetTimersResponse struct {
	Timers []Timer `json:"timers"`
}

var timers = []Timer{
	{Id: 1, Title: "Coding"},
	{Id: 2, Title: "Music Production"},
	{Id: 3, Title: "DJing"},
	{Id: 4, Title: "Piano"},
}

func GetTimers(c *gin.Context) {
	response := GetTimersResponse{Timers: timers}
	c.JSON(http.StatusOK, response)
}

func GetTimerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, t := range timers {
		if t.Id == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("timer with id=%d not found", id)})
}

func AddTimer(c *gin.Context) {
	var newTimer Timer

	if err := c.BindJSON(&newTimer); err != nil {
		return
	}

	timers = append(timers, newTimer)
	c.IndentedJSON(http.StatusCreated, newTimer)
}
