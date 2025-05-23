package migrations

import (
	"To_Do/pkg/logger"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"path/filepath"
)

func RunMigrations() error {
	migratPath, err := filepath.Abs("internal/migrations")
	if err != nil {
		logger.Logger.Error("ошибка получения пути к миграциям:", "err", err)
		return err
	}
	migrateURL := "file://" + filepath.ToSlash(migratPath)
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == " " {
		logger.Logger.Error("Ошибка получения URL базы данных", "err", err)
		return err
	}
	m, err := migrate.New(migrateURL, dbURL)
	if err != nil {
		logger.Logger.Error("Ошибка создании миграции", "err", err)
		return fmt.Errorf("ошибка инициализации мигратора: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Logger.Error("Ошибка применения миграций", "err", err)
		return err
	}
	logger.Logger.Info(" Миграции успешно применены!")
	return nil
}
