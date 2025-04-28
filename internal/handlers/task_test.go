package handlers

import (
	"To_Do/internal/models"
	"To_Do/internal/repository"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTaskHandler_Success(t *testing.T) {
	mockStorage := &repository.MockStorage{}

	mockStorage.On("CreateTask", "Test title", "Test content").Return(nil)

	body := models.CreateTaskRequest{
		Title:   "Test title",
		Content: "Test content",
	}

	JsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(JsonBody))
	rr := httptest.NewRecorder()
	handler := CreateTaskHandler(mockStorage)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "Задача добавлена", rr.Body.String())

	mockStorage.AssertExpectations(t)
}
