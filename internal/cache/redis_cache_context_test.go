package cache

import (
	"To_Do/internal/redis"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRedisCache_Set_Context(t *testing.T) {
	mockClient := NewMockRedis()
	cache := NewRedisCache(&redis.RedisClient{Client: mockClient, TTL: 5 * time.Second})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := cache.Set(ctx, "ctx_canceled_key", map[string]string{"foo": "bar"})
	require.Error(t, err, "ожидалась ошибка")
}
func TestRedisCache_Get_Context(t *testing.T) {
	mockClient := NewMockRedis()
	cache := NewRedisCache(&redis.RedisClient{Client: mockClient, TTL: 5 * time.Second})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var dest map[string]string
	err := cache.Get(ctx, "ctx_canceled_key", &dest)
	require.Error(t, err, "ожидалась ошибка")
}
func TestRedisCache_Del_ContextCanceled(t *testing.T) {
	mockClient := NewMockRedis()
	cache := NewRedisCache(&redis.RedisClient{
		Client: mockClient,
		TTL:    5 * time.Second,
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := cache.Del(ctx, "ctx_canceled_key")
	require.Error(t, err, "ожидалась ошибка при отменённом контексте")
}
