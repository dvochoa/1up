package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dvochoa/1up/db"
	"github.com/dvochoa/1up/models"
	test_db "github.com/dvochoa/1up/testdata/db"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Set up
	gin.SetMode(gin.TestMode)
	ctx := context.Background()

	// Set up test db
	_ = test_db.StartTestDatabase(ctx)
	defer test_db.StopTestDatabase()

	connStr, _ := test_db.GetTestDatabaseConnection(ctx)
	timerStore, _ := db.NewTimerStore(ctx, connStr)
	defer timerStore.CloseConnection(ctx)
	TimerStore = *timerStore

	// Run tests
	code := m.Run()

	// Teardown code goes here
	os.Exit(code)
}

func TestGetTimers(t *testing.T) {
	// Create a test context
	responseWriter := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseWriter)
	req, _ := http.NewRequest("GET", "/timers", nil)
	ctx.Request = req

	GetTimers(ctx)
	assert.EqualValues(t, http.StatusOK, responseWriter.Code)

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
	// Create a test context
	responseWriter := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseWriter)
	req, _ := http.NewRequest("GET", "/timers/", nil)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	GetTimerById(ctx)
	assert.EqualValues(t, http.StatusOK, responseWriter.Code)

	var result models.Timer
	err := json.Unmarshal(responseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, models.Timer{Id: 1, Title: "Coding"}, result)
}

func TestCreateTimer(t *testing.T) {
	// 1. No timer with Id=5 is found
	getResponseWriter := httptest.NewRecorder()
	getCtx, _ := gin.CreateTestContext(getResponseWriter)
	getReq, _ := http.NewRequest("GET", "/timers/", nil)
	getCtx.Request = getReq
	getCtx.Params = []gin.Param{{Key: "id", Value: "5"}}

	GetTimerById(getCtx)
	assert.EqualValues(t, http.StatusNotFound, getResponseWriter.Code)

	// 2. Create a timer
	postResponseWriter := httptest.NewRecorder()
	postCtx, _ := gin.CreateTestContext(postResponseWriter)

	newTimer := models.Timer{Id: 5, Title: "Cooking"}
	jsonValue, _ := json.Marshal(newTimer)

	postReq, _ := http.NewRequest(http.MethodPost, "/timers", bytes.NewBuffer(jsonValue))
	postReq.Header.Set("Content-Type", "application/json")
	postCtx.Request = postReq

	CreateTimer(postCtx)
	assert.EqualValues(t, http.StatusCreated, postResponseWriter.Code)

	var result models.Timer
	err := json.Unmarshal(postResponseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, models.Timer{Id: 5, Title: "Cooking"}, result)

	// 3. Timer with Id=5 is found
	getByIdResponseWriter := httptest.NewRecorder()
	getByIdCtx, _ := gin.CreateTestContext(getByIdResponseWriter)
	getByIdReq, _ := http.NewRequest("GET", "/timers/", nil)
	getByIdCtx.Request = getByIdReq
	getByIdCtx.Params = []gin.Param{{Key: "id", Value: "5"}}

	GetTimerById(getByIdCtx)
	assert.EqualValues(t, http.StatusOK, getByIdResponseWriter.Code)

	err = json.Unmarshal(getByIdResponseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, newTimer, result)
}
