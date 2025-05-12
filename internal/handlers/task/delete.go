package task

import (
	"To_Do/internal/models"
	"To_Do/internal/repository"
	"To_Do/pkg/logger"
	"encoding/json"
	"net/http"
)

func DeleteTaskHandler(s repository.StorageInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.DeleteTaskRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Не верное тело запроса", http.StatusBadRequest)
			return
		}

		err := s.DeleteTask(body.Id)
		if err != nil {
			logger.Logger.Error("Ошибка", "err", err)
			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Задача удалена"))
	}
}
