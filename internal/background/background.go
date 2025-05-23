package background

import (
	"To_Do/internal/cache"
	"To_Do/internal/repository"
	"To_Do/pkg/logger"
	"context"
	"time"
)

func StartTaskCacheRefresher(ctx context.Context, storage repository.StorageInterface, cache cache.Cache, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Logger.Info("Остановлено фоновое кеширование в Redis")
		case <-ticker.C:
			refreshTaskCache(ctx, storage, cache)
		}
	}
}

func refreshTaskCache(ctx context.Context, storage repository.StorageInterface, cache cache.Cache) {
	task, err := storage.ListTask()
	if err != nil {
		logger.Logger.Error("Ошибка получения задач из БД для кеша", "err", err)
		return
	}
	err = cache.Set(ctx, "task_all", task)
	if err != nil {
		logger.Logger.Error("Ошибка установки кеша задач", "err", err)
		return
	}

	logger.Logger.Info("Кеш задач успешно обновлён в фоне")
}
