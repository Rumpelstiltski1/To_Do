package handlers

import (
	"To_Do/internal/models"
	"To_Do/internal/repository"
	"To_Do/pkg/logger"
	"encoding/json"
	"net/http"
)

func PutTaskHandler(storage *repository.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodys models.PutTaskRequest
		if err := json.NewDecoder(r.Body).Decode(&bodys); err != nil {
			logger.Logger.Error("Неверное тело запроса", "err", err)
			http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		}
		err := storage.PutTask(bodys.Status, bodys.Id)
		if err != nil {
			logger.Logger.Error("Не удалось изменить задачу", "err", err)
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}
		response := models.PutTaskRequest{
			Id:     bodys.Id,
			Status: bodys.Status,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Logger.Error("Ошибка формирования ответа", "err", err)
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		}
	}
}
