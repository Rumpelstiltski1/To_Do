package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Commander interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Ping(ctx context.Context) *redis.StatusCmd
	Close() error
}
