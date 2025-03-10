package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dvochoa/1up/db"
	"github.com/dvochoa/1up/handlers"
	"github.com/dvochoa/1up/models"
	test_db "github.com/dvochoa/1up/testdata/db"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

// Runs before all tests
func TestMain(m *testing.M) {
	// Set up gin
	gin.SetMode(gin.TestMode)
	router = GetRouter()

	// Set up test db
	ctx := context.Background()
	_ = test_db.StartTestDatabase(ctx)
	defer test_db.StopTestDatabase()

	connStr, _ := test_db.GetTestDatabaseConnection(ctx)
	timerStore, _ := db.NewTimerStore(ctx, connStr)
	defer timerStore.CloseConnection(ctx)
	handlers.TimerStore = *timerStore

	// Run tests
	code := m.Run()

	// Teardown code goes here
	os.Exit(code)
}

func TestGetTimers(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/timers", nil)
	responseWriter := httptest.NewRecorder()
	router.ServeHTTP(responseWriter, req)

	assert.Equal(t, http.StatusOK, responseWriter.Code)

	var result models.GetTimersResponse
	err := json.Unmarshal(responseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	timers := result.Timers
	assert.Equal(t, models.Timer{Id: 1, Title: "Coding"}, timers[0])
	assert.Equal(t, models.Timer{Id: 2, Title: "Music Production"}, timers[1])
	assert.Equal(t, models.Timer{Id: 3, Title: "DJing"}, timers[2])
	assert.Equal(t, models.Timer{Id: 4, Title: "Piano"}, timers[3])
}

func TestGetTimerById(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/timers/1", nil)
	responseWriter := httptest.NewRecorder()
	router.ServeHTTP(responseWriter, req)

	assert.Equal(t, http.StatusOK, responseWriter.Code)

	var result models.Timer
	err := json.Unmarshal(responseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, models.Timer{Id: 1, Title: "Coding"}, result)
}

func TestCreateTimer(t *testing.T) {
	// 1. No timer with Id=5 is found
	getReq, _ := http.NewRequest(http.MethodGet, "/timers/5", nil)
	getResponseWriter := httptest.NewRecorder()
	router.ServeHTTP(getResponseWriter, getReq)

	assert.EqualValues(t, http.StatusNotFound, getResponseWriter.Code)

	// 2. Create a timer
	newTimer := models.Timer{Id: 5, Title: "Cooking"}
	jsonValue, _ := json.Marshal(newTimer)

	postReq, _ := http.NewRequest(http.MethodPost, "/timers", bytes.NewBuffer(jsonValue))
	postReq.Header.Set("Content-Type", "application/json")

	postResponseWriter := httptest.NewRecorder()
	router.ServeHTTP(postResponseWriter, postReq)

	assert.EqualValues(t, http.StatusCreated, postResponseWriter.Code)

	var result models.Timer
	err := json.Unmarshal(postResponseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.Equal(t, models.Timer{Id: 5, Title: "Cooking"}, result)

	// 3. Timer with Id=5 is found
	getByIdReq, _ := http.NewRequest(http.MethodGet, "/timers/5", nil)
	getByIdResponseWriter := httptest.NewRecorder()
	router.ServeHTTP(getByIdResponseWriter, getByIdReq)

	assert.EqualValues(t, http.StatusOK, getByIdResponseWriter.Code)

	err = json.Unmarshal(getByIdResponseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, newTimer, result)
}

func TestDeleteTimer(t *testing.T) {
	deleteReq, _ := http.NewRequest(http.MethodDelete, "/timers/1", nil)
	deleteResponseWriter := httptest.NewRecorder()
	router.ServeHTTP(deleteResponseWriter, deleteReq)

	assert.Equal(t, http.StatusNoContent, deleteResponseWriter.Code)

	getReq, _ := http.NewRequest(http.MethodGet, "/timers/1", nil)
	getResponseWriter := httptest.NewRecorder()
	router.ServeHTTP(getResponseWriter, getReq)

	assert.EqualValues(t, http.StatusNotFound, getResponseWriter.Code)
}
