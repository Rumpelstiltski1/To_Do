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

func TestPutTaskHandler_Success(t *testing.T) {
	mockStorage := &repository.MockStorage{}

	mockStorage.On("PutTask", true, 1).Return(nil)

	task := models.PutTaskRequest{
		Status: true,
		Id:     1,
	}
	jsonBody, _ := json.Marshal(task)

	req := httptest.NewRequest(http.MethodPut, "/done", bytes.NewReader(jsonBody))
	rr := httptest.NewRecorder()

	handlers := PutTaskHandler(mockStorage)
	handlers.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resTask models.PutTaskRequest
	err := json.Unmarshal(rr.Body.Bytes(), &resTask)
	assert.NoError(t, err)

	assert.Equal(t, task.Id, resTask.Id)
	assert.Equal(t, task.Status, resTask.Status)

	mockStorage.AssertExpectations(t)
}
