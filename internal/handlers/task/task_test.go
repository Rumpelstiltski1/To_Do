package task

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
	var id = 123
	mockStorage.On("CreateTask", "Test title", "Test content").Return(id, nil)

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
	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Задача добавлена. ID созданной задачи:", response["message"])
	assert.EqualValues(t, id, response["id"])
	mockStorage.AssertExpectations(t)
}
