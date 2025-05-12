package task

import (
	"To_Do/internal/repository"
	"To_Do/pkg/logger"
	"encoding/json"
	"net/http"
)

func ListTaskHandler(storage repository.StorageInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		list, err := storage.ListTask()
		if err != nil {
			logger.Logger.Error("Ошибка выполнения фyнкции", "err", err)
			http.Error(w, "Request execution error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(list); err != nil {
			logger.Logger.Error("Ошибка обработки тело ответа", "err", err)
			http.Error(w, "Не верное тело ответа", http.StatusInternalServerError)
			return
		}
	}
}
