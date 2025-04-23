package repository

import "database/sql"

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateTask(title string, content string) error {
	_, err := s.db.Exec("INSERT INTO tasks(title, content) VALUES ($1, $2)", title, content)
	return err
}
