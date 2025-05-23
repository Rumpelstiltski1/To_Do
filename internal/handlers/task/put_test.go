package task

import (
	"To_Do/internal/cache"
	"To_Do/internal/models"
	"To_Do/internal/redis"
	"To_Do/internal/repository"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPutTaskHandler_Success(t *testing.T) {
	mockStorage := &repository.MockStorage{}

	mockStorage.On("PutTask", true, 1).Return(nil)

	task := models.PutTaskRequest{
		Status: true,
		Id:     1,
	}
	jsonBody, _ := json.Marshal(task)

	mockRedis := cache.NewMockRedis()
	mockCache := cache.NewRedisCache(&redis.RedisClient{
		Client: mockRedis,
		TTL:    time.Second * 5,
	})

	mockRedis.Data["task:5"] = []byte("some data")
	mockRedis.Data["task_all"] = []byte("cached list")

	req := httptest.NewRequest(http.MethodPut, "/done", bytes.NewReader(jsonBody))
	rr := httptest.NewRecorder()

	handlers := PutTaskHandler(mockStorage, mockCache)
	handlers.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resTask models.PutTaskRequest
	err := json.Unmarshal(rr.Body.Bytes(), &resTask)
	assert.NoError(t, err)

	assert.Equal(t, task.Id, resTask.Id)
	assert.Equal(t, task.Status, resTask.Status)

	_, exists := mockRedis.Data["task:5"]
	assert.False(t, exists, "task:5 должен быть удален")
	_, exists = mockRedis.Data["task_all"]
	assert.False(t, exists, "task_all должен быть удален")

	mockStorage.AssertExpectations(t)
}
