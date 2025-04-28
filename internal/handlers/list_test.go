package handlers

import (
	"To_Do/internal/models"
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
		Id:         1,
		Title:      "Test task",
		Content:    "Test content",
		Status:     false,
		Created_at: time.Now(),
	}, {
		Id:         2,
		Title:      "Test task2",
		Content:    "Test content2",
		Status:     true,
		Created_at: time.Now(),
	},
	}

	mockStorage.On("ListTask").Return(tasks, nil)

	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	rr := httptest.NewRecorder()

	handlers := ListTaskHandler(mockStorage)
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
	mockStorage.AssertExpectations(t)
}
