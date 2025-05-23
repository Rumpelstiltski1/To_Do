package repository

import (
	"To_Do/internal/models"
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateTask(title, content string) (*models.ListTaskResponse, error) {
	var task models.ListTaskResponse
	err := s.db.QueryRow("INSERT INTO tasks(title, content) VALUES ($1, $2) RETURNING id, title, content, status, created_at", title, content).Scan(&task.Id, &task.Title, &task.Content, &task.Status, &task.CreatedAt)

	return &task, err
}

func (s *Storage) PutTask(status bool, id int) error {
	result, err := s.db.Exec("update tasks set status= $1 WHERE id=$2", status, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось получить количество изменённых строк: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	return nil
}
func (s *Storage) ListTask() ([]models.ListTaskResponse, error) {
	tasks := []models.ListTaskResponse{}
	rows, err := s.db.Query("select id, title, content, status, created_at from tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		task := models.ListTaskResponse{}
		err := rows.Scan(&task.Id, &task.Title, &task.Content, &task.Status, &task.CreatedAt)
		if err != nil {
			return nil, err

		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *Storage) DeleteTask(id int) error {
	result, err := s.db.Exec("delete from tasks where id=$1", id)
	if err != nil {
		return err
	}
	rowsAffectd, _ := result.RowsAffected()
	if rowsAffectd == 0 {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	return nil
}
