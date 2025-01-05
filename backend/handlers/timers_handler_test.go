package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandlerReturnsTimers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	responseWriter := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseWriter)

	TimersHandler(ctx)

	assert.EqualValues(t, http.StatusOK, responseWriter.Code)

	var result TimerResponse
	json.Unmarshal(responseWriter.Body.Bytes(), &result)
	timers = result.Timers

	assert.Equal(t, timers[0], Timer{Id: 1, Title: "Coding"})
	assert.Equal(t, timers[1], Timer{Id: 2, Title: "Music Production"})
	assert.Equal(t, timers[2], Timer{Id: 3, Title: "DJing"})
	assert.Equal(t, timers[3], Timer{Id: 4, Title: "Piano"})
}
