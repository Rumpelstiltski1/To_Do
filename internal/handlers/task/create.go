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
	"sync"
)

func CreateTaskHandler(storage repository.StorageInterface, cache cache.Cache, producer kafka.KafkaProducer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.CreateTaskRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
			return
		}

		task, err := storage.CreateTask(body.Title, body.Content)
		if err != nil {
			logger.Logger.Error("Ошибка при создании задачи:", "err", err)
			http.Error(w, "Failed to create task", http.StatusInternalServerError)

			return
		}

		ctx := r.Context()
		cacheKey := fmt.Sprint("task:", task.Id)
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			if err := cache.Set(ctx, cacheKey, task); err != nil {
				logger.Logger.Error("Ошибка кеширования задачи", "err", err)
			}
		}()

		go func() {
			defer wg.Done()
			if err := cache.Del(ctx, "task_all"); err != nil {
				logger.Logger.Error("Не удалось очистить кеш списка задач:", "err", err)
			}
		}()

		wg.Wait()

		event := fmt.Sprintf("action=create, id=%d, title=%q, content=%q, time=%s", task.Id, task.Title, task.Content, task.CreatedAt)
		if err := producer.SendMessage(ctx, []byte("create"), []byte(event)); err != nil {
			logger.Logger.Error("Не удалось отправить сообщение в Kafka", "err", err)
		}
		response := map[string]interface{}{
			"message": "Задача добавлена. ID созданной задачи:",
			"id":      task.Id,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
