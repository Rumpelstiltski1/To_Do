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

func TestCreateTaskHandler_Success_WithCache(t *testing.T) {
	mockStorage := &repository.MockStorage{}
	task := &models.ListTaskResponse{
		Id:      123,
		Title:   "Test title",
		Content: "Test content",
		Status:  false,
	}
	mockStorage.On("CreateTask", "Test title", "Test content").Return(task, nil)

	body := models.CreateTaskRequest{
		Title:   "Test title",
		Content: "Test content",
	}

	JsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(JsonBody))
	rr := httptest.NewRecorder()

	mockRedis := cache.NewMockRedis()
	mockCache := cache.NewRedisCache(&redis.RedisClient{
		Client: mockRedis,
		TTL:    time.Second * 5,
	})
	handler := CreateTaskHandler(mockStorage, mockCache)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Задача добавлена. ID созданной задачи:", response["message"])
	assert.EqualValues(t, task.Id, response["id"])

	// Проверим, что task закешировался
	_, exists := mockRedis.Data["task:123"]
	assert.True(t, exists)

	// Проверим, что task_all был удалён
	_, deleted := mockRedis.Data["task_all"]
	assert.False(t, deleted) // был удалён — поэтому его нет

	mockStorage.AssertExpectations(t)
}
