package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "fake_password" // TODO: Change this to real password but don't commit with it
	dbname   = "oneup"
)

// TODO: What is the json here? Its called a struct tag?
type Timer struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type GetTimersResponse struct {
	Timers []Timer `json:"timers"`
}

var fakeTimers = []Timer{
	{Id: 1, Title: "Coding"},
	{Id: 2, Title: "Music Production"},
	{Id: 3, Title: "DJing"},
	{Id: 4, Title: "Piano"},
}

func GetTimers(c *gin.Context) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname))
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), "select * from timers")
	timers, err := pgx.CollectRows(rows, pgx.RowToStructByName[Timer])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		os.Exit(1)
	}

	response := GetTimersResponse{Timers: timers}
	c.JSON(http.StatusOK, response)
}

func GetTimerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, t := range fakeTimers {
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

	fakeTimers = append(fakeTimers, newTimer)
	c.IndentedJSON(http.StatusCreated, newTimer)
}
