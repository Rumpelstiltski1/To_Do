package httpserver

import (
	"To_Do/internal/cache"
	"To_Do/internal/handlers/task"
	"To_Do/internal/kafka"
	"To_Do/internal/metricks"
	"To_Do/internal/repository"
	Mymiddleware "To_Do/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func NewRouter(storage *repository.Storage, cache cache.Cache, producer kafka.KafkaProducer) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(Mymiddleware.SlogMiddleware)
	router.Use(metricks.MetricsMiddleware)

	router.Post("/create", task.CreateTaskHandler(storage, cache, producer))
	router.Get("/list", task.ListTaskHandler(storage, cache))
	router.Delete("/delete", task.DeleteTaskHandler(storage, cache, producer))
	router.Put("/done", task.PutTaskHandler(storage, cache, producer))
	router.Handle("/metrics", promhttp.Handler())
	return router
}
