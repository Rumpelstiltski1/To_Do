package cache

import (
	"To_Do/internal/redis"
	"To_Do/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
)

type RedisCache struct {
	client *redis.RedisClient
}

type Cache interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string, dest any) error
	Del(ctx context.Context, key string) error
}

func NewRedisCache(client *redis.RedisClient) *RedisCache {
	return &RedisCache{client: client}
}

func (c *RedisCache) Set(ctx context.Context, key string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		logger.Logger.Error("Ошибка преобразования значения в JSON", "err", err)
		return fmt.Errorf("cache: failed to marshal: %w", err)
	}
	return c.client.Client.Set(ctx, key, data, c.client.TTL).Err()
}

func (c *RedisCache) Get(ctx context.Context, key string, dest any) error {
	data, err := c.client.Client.Get(ctx, key).Bytes()
	if err != nil {
		logger.Logger.Error("", "err", err)
		return err
	}
	return json.Unmarshal(data, dest)
}
func (c *RedisCache) Del(ctx context.Context, key string) error {
	return c.client.Client.Del(ctx, key).Err()
}
