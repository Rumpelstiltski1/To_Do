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

func PutTaskHandler(storage repository.StorageInterface, cache cache.Cache, producer kafka.KafkaProducer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.PutTaskRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			logger.Logger.Error("Неверное тело запроса", "err", err)
			http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
			return
		}
		err := storage.PutTask(body.Status, body.Id)
		if err != nil {
			logger.Logger.Error("Не удалось изменить задачу", "err", err)
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}
		response := models.PutTaskRequest{
			Id:     body.Id,
			Status: body.Status,
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

		event := fmt.Sprintf("action=delete, id=%d, status=%t", body.Id, body.Status)
		if err := producer.SendMessage(ctx, []byte("delete"), []byte(event)); err != nil {
			logger.Logger.Error("Не удалось отправить сообщение в Kafka", "err", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Logger.Error("Ошибка формирования ответа", "err", err)
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		}
	}
}
