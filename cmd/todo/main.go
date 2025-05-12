package main

import (
	"To_Do/config"
	"To_Do/internal/httpserver"
	"To_Do/internal/migrations"
	"To_Do/internal/repository"
	"To_Do/pkg/database"
	"To_Do/pkg/logger"
	"context"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	if os.Getenv("ENV") != "production" {
		if err := migrations.RunMigrations(); err != nil {
			logger.Logger.Error("Ошибка запуска миграций ", "err", err)
			log.Fatal(err)
		}
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Ошибка при загрузке .env файла в dev-среде")

		}
	}
	cfg := config.LoadConfig()

	logger.InitLog(cfg.LogLevel)
	defer logger.CloseFile()

	db, err := database.InitDb(cfg.Database)
	if err != nil {
		logger.Logger.Error("Ошибка инициализации данных", "err", err)
		logger.CloseFile()
		os.Exit(1)
	}
	defer db.Close()

	storage := repository.NewStorage(db)
	server := httpserver.NewServer(cfg, storage)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Logger.Info("Сервер запущен", "port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Error("Ошибка сервера", "err", err)
			cancel()
		}

	}()

	<-sigChan
	logger.Logger.Info("Начало завершения работы")
	cancel()

	if err := httpserver.ShutdownServer(ctx, server); err != nil {
		logger.Logger.Error("Не удалось завершить работу сервера за отведенное время", "err", err)
	} else {
		logger.Logger.Info("Работа сервера прекращена корректно")
	}
	wg.Wait()
}
