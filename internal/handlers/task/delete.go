package task

import (
	"To_Do/internal/cache"
	"To_Do/internal/kafka"
	"To_Do/internal/models"
	"To_Do/internal/repository"
	"To_Do/pkg/logger"
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteTaskHandler(s repository.StorageInterface, cache cache.Cache, producer kafka.KafkaProducer) http.HandlerFunc {
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

		ctx := r.Context()
		err = cache.Del(ctx, fmt.Sprintf("task:%d", body.Id))
		if err != nil {
			logger.Logger.Error("Ошибка удаления кеша задачи", "err", err)
		}
		err = cache.Del(ctx, "task_all")
		if err != nil {
			logger.Logger.Error("Не удалось очистить кеш", "err", err)
		}
		event := fmt.Sprintf("action=update, id=%d", body.Id)
		if err := producer.SendMessage(ctx, []byte("delete"), []byte(event)); err != nil {
			logger.Logger.Error("Не удалось отправить сообщение в Kafka", "err", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Задача удалена"))
	}
}
