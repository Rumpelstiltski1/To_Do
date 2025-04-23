package handlers

import (
	"To_Do/internal/models"
	"To_Do/internal/repository"
	"encoding/json"
	"net/http"
)

func CreateTaskHandler(storage *repository.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.CreateTaskRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Не верное тело запроса", http.StatusBadRequest)
			return
		}

		err := storage.CreateTask(body.Title, body.Content)
		if err != nil {
			http.Error(w, "Failed to create task", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Задача создана"))
	}
}
