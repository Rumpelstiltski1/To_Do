package repository

import "To_Do/internal/models"

type StorageInterface interface {
	CreateTask(title, content string) error
	ListTask() ([]models.ListTaskResponse, error)
	DeleteTask(id int) error
	PutTask(status bool, id int) error
}
