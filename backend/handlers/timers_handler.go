package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Timer struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type TimerResponse struct {
	Timers []Timer `json:"timers"`
}

var timers = []Timer{
	{Id: 1, Title: "Coding"},
	{Id: 2, Title: "Music Production"},
	{Id: 3, Title: "DJing"},
	{Id: 4, Title: "Piano"},
}

func TimersHandler(c *gin.Context) {
	response := TimerResponse{Timers: timers}
	c.JSON(http.StatusOK, response)
}
