package httpserver

import (
	"To_Do/config"
	"To_Do/internal/cache"
	"To_Do/internal/kafka"
	"To_Do/internal/repository"
	"context"
	"net/http"
	"time"
)

func NewServer(cfg *config.Config, storage *repository.Storage, cache cache.Cache, producer kafka.KafkaProducer) *http.Server {
	server := &http.Server{
		Addr:              cfg.Port,
		Handler:           NewRouter(storage, cache, producer),
		ReadTimeout:       cfg.Server.ReadTimeout,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
		WriteTimeout:      cfg.Server.WriteTimeout,
		IdleTimeout:       cfg.Server.IdleTimeout,
	}
	return server
}

func ShutdownServer(ctx context.Context, server *http.Server) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return server.Shutdown(shutdownCtx)
}
