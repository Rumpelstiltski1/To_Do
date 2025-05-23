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

func TestDeleteTaskHandler(t *testing.T) {
	mockStorage := &repository.MockStorage{}
	mockStorage.On("DeleteTask", 5).Return(nil)

	tasks := &models.DeleteTaskRequest{
		Id: 5,
	}
	jsonBody, _ := json.Marshal(tasks)

	mockRedis := cache.NewMockRedis()
	mockCache := cache.NewRedisCache(&redis.RedisClient{
		Client: mockRedis,
		TTL:    time.Second * 5,
	})

	mockRedis.Data["task:5"] = []byte("some data")
	mockRedis.Data["task_all"] = []byte("cached list")

	req := httptest.NewRequest(http.MethodDelete, "/delete", bytes.NewReader(jsonBody))
	rr := httptest.NewRecorder()
	handlres := DeleteTaskHandler(mockStorage, mockCache)
	handlres.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Задача удалена", rr.Body.String())

	_, exists := mockRedis.Data["task:5"]
	assert.False(t, exists, "task:5 должен быть удален")
	_, exists = mockRedis.Data["task_all"]
	assert.False(t, exists, "task_all должен быть удален")

	mockStorage.AssertExpectations(t)
}
