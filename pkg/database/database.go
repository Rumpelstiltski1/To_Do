package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func InitDb(url string) (*sql.DB, error) {
	connectURL := url
	db, err := sql.Open("postgres", connectURL) // Подключение к бд
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к базе данных: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Не удалось проверить соединение:%w", err)
	}
	return db, nil
}
