package redis

import (
	"To_Do/config"
	"To_Do/pkg/redisx"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	Client redisx.Commander
	TTL    time.Duration
}

func NewRedisClient(cfg *config.Config) *RedisClient {
	realClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	return &RedisClient{Client: realClient, TTL: cfg.RedisTTL}
}

func (r *RedisClient) Ping(ctx context.Context) error {
	_, err := r.Client.Ping(ctx).Result()
	return err
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}
