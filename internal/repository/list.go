package repository

import "To_Do/internal/models"

func (s *Storage) ListTask() ([]models.ListTaskResponse, error) {
	tasks := []models.ListTaskResponse{}
	rows, err := s.db.Query("select id, title, content, status, created_at from tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		task := models.ListTaskResponse{}
		err := rows.Scan(&task.Id, &task.Title, &task.Content, &task.Status, &task.Created_at)
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
