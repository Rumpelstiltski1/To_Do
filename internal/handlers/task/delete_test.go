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

func TestDeleteTaskHandler(t *testing.T) {
	mockStorage := &repository.MockStorage{}
	mockStorage.On("DeleteTask", 5).Return(nil)

	tasks := models.DeleteTaskRequest{
		Id: 5,
	}
	jsonBody, _ := json.Marshal(tasks)

	req := httptest.NewRequest(http.MethodDelete, "/delete", bytes.NewReader(jsonBody))
	rr := httptest.NewRecorder()
	handlres := DeleteTaskHandler(mockStorage)
	handlres.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Задача удалена", rr.Body.String())

	mockStorage.AssertExpectations(t)
}
