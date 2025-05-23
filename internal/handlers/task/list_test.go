package task

import (
	"To_Do/internal/cache"
	"To_Do/internal/models"
	"To_Do/internal/redis"
	"To_Do/internal/repository"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestListTaskHandler_Success(t *testing.T) {
	mockStorage := &repository.MockStorage{}

	tasks := []models.ListTaskResponse{{
		Id:        1,
		Title:     "Test task",
		Content:   "Test content",
		Status:    false,
		CreatedAt: time.Now(),
	}, {
		Id:        2,
		Title:     "Test task2",
		Content:   "Test content2",
		Status:    true,
		CreatedAt: time.Now(),
	},
	}

	mockStorage.On("ListTask").Return(tasks, nil)

	mockRedis := cache.NewMockRedis()
	mockCache := cache.NewRedisCache(&redis.RedisClient{
		Client: mockRedis,
		TTL:    time.Second * 5,
	})

	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	rr := httptest.NewRecorder()

	mockRedis.Data["task:5"] = []byte("some data")
	mockRedis.Data["task_all"] = []byte("cached list")

	handlers := ListTaskHandler(mockStorage, mockCache)
	handlers.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var actualTask []models.ListTaskResponse
	err := json.Unmarshal(rr.Body.Bytes(), &actualTask)
	assert.NoError(t, err)

	for i, _ := range tasks {
		assert.Equal(t, tasks[i].Id, actualTask[i].Id)
		assert.Equal(t, tasks[i].Title, actualTask[i].Title)
		assert.Equal(t, tasks[i].Content, actualTask[i].Content)
		assert.Equal(t, tasks[i].Status, actualTask[i].Status)
	}

	cached, ok := mockRedis.Data["task_all"]
	assert.True(t, ok, "task_all должен быть записан в кэш")
	assert.NotEmpty(t, cached, "task_all не должен быть пустым")
	mockStorage.AssertExpectations(t)
}
