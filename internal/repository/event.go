package repository

import (
	"To_Do/internal/models"
	"context"
)

func (s *Storage) SaveEvent(ctx context.Context, event models.TaskEvent) error {
	query := `INSERT INTO task_events (task_id, action, title, content, status, created_at)
          VALUES ($1, $2, $3, $4, $5, now())`
	_, err := s.db.ExecContext(ctx, query, event.TaskID, event.Action, event.Title, event.Content, event.Status, event.CreatedAt)
	return err
}
