package httpserver

import (
	"To_Do/internal/handlers/task"
	"To_Do/internal/repository"
	Mymiddleware "To_Do/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func NewRouter(storage *repository.Storage) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(Mymiddleware.SlogMiddleware)

	router.Post("/create", task.CreateTaskHandler(storage))
	router.Get("/list", task.ListTaskHandler(storage))
	router.Delete("/delete", task.DeleteTaskHandler(storage))
	router.Put("/done", task.PutTaskHandler(storage))
	return router
}
