package task

import (
	"To_Do/internal/models"
	"To_Do/internal/repository"
	"To_Do/pkg/logger"
	"encoding/json"
	"net/http"
)

func CreateTaskHandler(storage repository.StorageInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.CreateTaskRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Не верное тело запроса", http.StatusBadRequest)
			return
		}

		id, err := storage.CreateTask(body.Title, body.Content)
		if err != nil {
			logger.Logger.Error("Ошибка при создании задачи:", "err", err)
			http.Error(w, "Failed to create task", http.StatusInternalServerError)

			return
		}

		defent := map[string]interface{}{
			"message": "Задача добавлена. ID созданной задачи:",
			"id":      id,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(defent)
	}
}
