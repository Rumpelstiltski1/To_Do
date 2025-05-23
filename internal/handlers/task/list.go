package task

import (
	"To_Do/internal/cache"
	"To_Do/internal/models"
	"To_Do/internal/repository"
	"To_Do/pkg/logger"
	"encoding/json"
	"net/http"
)

func ListTaskHandler(storage repository.StorageInterface, cache cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var task []models.ListTaskResponse

		ctx := r.Context()
		const cacheKey = "task_all"
		err := cache.Get(ctx, cacheKey, &task)
		if err == nil {
			logger.Logger.Info("Кэш задачи найден")

			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(task)
			return
		}

		task, err = storage.ListTask()
		if err != nil {
			logger.Logger.Error("Ошибка выполнения фyнкции", "err", err)
			http.Error(w, "Request execution error", http.StatusInternalServerError)
			return
		}

		if err := cache.Set(ctx, cacheKey, task); err != nil {
			logger.Logger.Error("Ошибка при сохранении списка задач в Redis", "err", err)
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(task); err != nil {
			logger.Logger.Error("Ошибка обработки тело ответа", "err", err)
			http.Error(w, "Не верное тело ответа", http.StatusInternalServerError)
			return
		}
	}
}
