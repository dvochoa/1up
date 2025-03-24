package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

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
	timerStore, err := db.NewTimerStore(ctx, connStr)
	if err != nil {
		fmt.Println(err)
	}
	defer timerStore.CloseConnection(ctx)
	handlers.TimerStore = *timerStore

	// Run tests
	code := m.Run()

	// Teardown code goes here
	os.Exit(code)
}

func TestGetTimers(t *testing.T) {
	responseWriter := serveHTTP(http.MethodGet, "/users/1/timers", nil)
	assert.Equal(t, http.StatusOK, responseWriter.Code)

	var result models.GetTimersResponse
	err := json.Unmarshal(responseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	timers := result.Timers
	assert.Equal(t, 4, len(timers))
	assert.Equal(t, models.Timer{Id: 1, OwnerId: 1, Title: "Coding", TotalTime: 5400}, timers[0])
	assert.Equal(t, models.Timer{Id: 2, OwnerId: 1, Title: "Music Production", TotalTime: 2700}, timers[1])
	assert.Equal(t, models.Timer{Id: 3, OwnerId: 1, Title: "DJing", TotalTime: 600}, timers[2])
	assert.Equal(t, models.Timer{Id: 4, OwnerId: 1, Title: "Piano", TotalTime: 0}, timers[3])
}

func TestGetTimerHistory(t *testing.T) {
	responseWriter := serveHTTP(http.MethodGet, "/timers/1", nil)
	assert.Equal(t, http.StatusOK, responseWriter.Code)

	var result models.GetTimerHistoryResponse
	err := json.Unmarshal(responseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	timerSessions := result.TimerSessions

	// Convert timestamps to UTC so that timestamp value is constant no matter the timezone this test is ran in
	for i := range timerSessions {
		timerSessions[i].SessionTimestamp = timerSessions[i].SessionTimestamp.UTC()
	}
	assert.Equal(t, models.TimerSession{Id: 1, SessionTimestamp: time.Date(2025, 3, 15, 12, 0, 0, 0, time.UTC), SessionDurationInSeconds: 3600}, timerSessions[0])
	assert.Equal(t, models.TimerSession{Id: 2, SessionTimestamp: time.Date(2025, 3, 20, 10, 0, 0, 0, time.UTC), SessionDurationInSeconds: 1800}, timerSessions[1])
}

// TODO: When calling createTimer the frontend won't be able to supply the id, any way to make the model have id nullable?
// Or should I create a different interface?
func TestCreateTimer(t *testing.T) {
	// 1. No timer with Id=6 is found
	getResponseWriter := serveHTTP(http.MethodGet, "/users/2/timers", nil)
	assert.EqualValues(t, http.StatusOK, getResponseWriter.Code)

	var getTimersResult models.GetTimersResponse
	err := json.Unmarshal(getResponseWriter.Body.Bytes(), &getTimersResult)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(getTimersResult.Timers))
	assert.Equal(t, int64(5), getTimersResult.Timers[0].Id)

	// 2. Create a timer
	newTimer := models.Timer{Id: 6, OwnerId: 2, Title: "Cooking"}
	jsonValue, _ := json.Marshal(newTimer)

	postResponseWriter := serveHTTP(http.MethodPost, "/timers", bytes.NewBuffer(jsonValue))
	assert.EqualValues(t, http.StatusCreated, postResponseWriter.Code)

	var result models.Timer
	err = json.Unmarshal(postResponseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.Equal(t, newTimer, result)

	// 3. Timer with Id=6 is found
	getResponseWriter = serveHTTP(http.MethodGet, "/users/2/timers", nil)
	assert.EqualValues(t, http.StatusOK, getResponseWriter.Code)

	err = json.Unmarshal(getResponseWriter.Body.Bytes(), &getTimersResult)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(getTimersResult.Timers))
	assert.Equal(t, newTimer, getTimersResult.Timers[1])
}

func TestUpdateTimerSettings(t *testing.T) {
	// 1. Update timer
	updatedTimer := models.Timer{Id: 1, OwnerId: 1, Title: "Dancing"}
	jsonValue, _ := json.Marshal(updatedTimer)

	putResponseWriter := serveHTTP(http.MethodPut, "/timers/1", bytes.NewBuffer(jsonValue))
	assert.EqualValues(t, http.StatusNoContent, putResponseWriter.Code)

	// 2. Get timers
	getResponseWriter := serveHTTP(http.MethodGet, "/users/1/timers", nil)
	assert.EqualValues(t, http.StatusOK, getResponseWriter.Code)

	var getTimersResult models.GetTimersResponse
	err := json.Unmarshal(getResponseWriter.Body.Bytes(), &getTimersResult)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(getTimersResult.Timers))
	assert.Equal(t, models.Timer{Id: 1, OwnerId: 1, Title: "Dancing", TotalTime: 5400}, getTimersResult.Timers[0])
}

func TestDeleteTimer(t *testing.T) {
	deleteResponseWriter := serveHTTP(http.MethodDelete, "/timers/1", nil)
	assert.Equal(t, http.StatusNoContent, deleteResponseWriter.Code)

	getResponseWriter := serveHTTP(http.MethodGet, "/users/1/timers", nil)
	assert.EqualValues(t, http.StatusOK, getResponseWriter.Code)

	var getTimersResult models.GetTimersResponse
	err := json.Unmarshal(getResponseWriter.Body.Bytes(), &getTimersResult)
	assert.Nil(t, err)

	for _, timer := range getTimersResult.Timers {
		assert.NotEqual(t, 1, timer.Id)
	}
}

func serveHTTP(httpMethod string, url string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(httpMethod, url, body)

	switch httpMethod {
	case http.MethodPut, http.MethodPost:
		req.Header.Set("Content-Type", "application/json")
	}

	responseWriter := httptest.NewRecorder()
	router.ServeHTTP(responseWriter, req)

	return responseWriter
}
