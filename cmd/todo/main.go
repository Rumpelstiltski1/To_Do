package main

import (
	"To_Do/config"
	"To_Do/pkg/database"
	"To_Do/pkg/logger"
	Mymiddleware "To_Do/pkg/middleware"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env")
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

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(Mymiddleware.SlogMiddleware)

	server := &http.Server{
		Addr:              cfg.Port,
		Handler:           router,
		ReadTimeout:       cfg.Server.ReadTimeout,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
		WriteTimeout:      cfg.Server.WriteTimeout,
		IdleTimeout:       cfg.Server.IdleTimeout,
	}

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

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Error("Не удалось завершить работу сервера за отведенное время", "err", err)
	} else {
		logger.Logger.Info("Работа сервера прекращена корректно")
	}
	wg.Wait()
}
