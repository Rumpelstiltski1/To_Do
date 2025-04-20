package main

import (
	"To_Do/config"
	"To_Do/pkg/database"
	"To_Do/pkg/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// TODO: init config: os.Getenv +
	// TODO: init logger: log/slog +
	// TODO: init storage: PostgreSQL +
	// TODO: init router: chi
	// TODO: run server

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env")
	}

	cfg := config.LoadConfig()
	logger.InitLog(cfg.LogLevel)
	defer logger.CloseFile()
	logger.Logger.Info("Запуск приложения")

	db, err := database.InitDb(cfg.Database)
	if err != nil {
		logger.Logger.Error("Ошибка инициализации данных", "err", err)
		logger.CloseFile()
		os.Exit(1)
	}

	logger.Logger.Info("Конфигурация загружена",
		"port", cfg.Port,
		"env", cfg.Env,
		"log_level", cfg.LogLevel,
	)

	logger.Logger.Info("Успешное подключение к базе данных")
	// TODO: передать db в слой хранилища (repository/service)
	_ = db
	logger.Logger.Info("Приложение запущено и готово к работе")

}
