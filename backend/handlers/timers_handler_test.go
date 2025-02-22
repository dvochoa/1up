package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Setup code goes here
	gin.SetMode(gin.TestMode)
	code := m.Run()

	// Teardown code goes here
	os.Exit(code)
}

func TestGetTimers(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseWriter)

	GetTimers(ctx)
	assert.EqualValues(t, http.StatusOK, responseWriter.Code)

	var result GetTimersResponse
	err := json.Unmarshal(responseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	timers = result.Timers
	assert.Equal(t, Timer{Id: 1, Title: "Coding"}, timers[0])
	assert.Equal(t, Timer{Id: 2, Title: "Music Production"}, timers[1])
	assert.Equal(t, Timer{Id: 3, Title: "DJing"}, timers[2])
	assert.Equal(t, Timer{Id: 4, Title: "Piano"}, timers[3])
}

func TestGetTimerById(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseWriter)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	GetTimerById(ctx)
	assert.EqualValues(t, http.StatusOK, responseWriter.Code)

	var result Timer
	err := json.Unmarshal(responseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, Timer{Id: 1, Title: "Coding"}, result)
}

func TestAddTimer(t *testing.T) {
	// 1. No timer with Id=5 is found
	getResponseWriter := httptest.NewRecorder()
	getCtx, _ := gin.CreateTestContext(getResponseWriter)
	getCtx.Params = []gin.Param{{Key: "id", Value: "5"}}
	GetTimerById(getCtx)
	assert.EqualValues(t, http.StatusNotFound, getResponseWriter.Code)

	// 2. Create a timer
	postResponseWriter := httptest.NewRecorder()
	postCtx, _ := gin.CreateTestContext(postResponseWriter)

	newTimer := Timer{Id: 5, Title: "Cooking"}
	jsonValue, _ := json.Marshal(newTimer)

	request, _ := http.NewRequest(http.MethodPost, "/timers", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")

	postCtx.Request = request

	AddTimer(postCtx)
	assert.EqualValues(t, http.StatusCreated, postResponseWriter.Code)

	var result Timer
	err := json.Unmarshal(postResponseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, Timer{Id: 5, Title: "Cooking"}, result)

	// 3. Timer with Id=5 is found
	getByIdResponseWriter := httptest.NewRecorder()
	getByIdCtx, _ := gin.CreateTestContext(getByIdResponseWriter)
	getByIdCtx.Params = []gin.Param{{Key: "id", Value: "5"}}
	GetTimerById(getByIdCtx)
	assert.EqualValues(t, http.StatusOK, getByIdResponseWriter.Code)

	err = json.Unmarshal(getByIdResponseWriter.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, newTimer, result)
}
