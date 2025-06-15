package repository

import (
	"To_Do/internal/models"
	"context"
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) CreateTask(title, content string) (*models.ListTaskResponse, error) {
	args := m.Called(title, content)
	return args.Get(0).(*models.ListTaskResponse), args.Error(1)
}

func (m *MockStorage) ListTask() ([]models.ListTaskResponse, error) {
	args := m.Called()
	return args.Get(0).([]models.ListTaskResponse), args.Error(1)
}

func (m *MockStorage) DeleteTask(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockStorage) PutTask(status bool, id int) error {
	args := m.Called(status, id)
	return args.Error(0)
}
func (m *MockStorage) SaveEvent(ctx context.Context, event models.TaskEvent) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}
