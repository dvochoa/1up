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
	password = "fake_password"
	dbname   = "oneup"
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
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname))
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select id, title from timers")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string

		err = rows.Scan(&id, &title)
		if err != nil {
			fmt.Printf("Scan error: %v", err)
			return
		}

		fmt.Println(id, title)
	}

	// The first error encountered by the original Query call, rows.Next or rows.Scan will be returned here.
	if rows.Err() != nil {
		fmt.Printf("rows error: %v", rows.Err())
		return
	}

	// Still returns stubbed response for now
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
