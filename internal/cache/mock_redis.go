package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

type MockRedisClient struct {
	Data map[string][]byte
	mu   sync.RWMutex
}

func NewMockRedis() *MockRedisClient {
	return &MockRedisClient{
		Data: make(map[string][]byte),
	}
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	m.mu.Lock()
	defer m.mu.Unlock()

	bytes, ok := value.([]byte)
	if !ok {
		cmd := redis.NewStatusCmd(ctx)
		cmd.SetErr(redis.Nil)
		return cmd
	}
	m.Data[key] = bytes

	cmd := redis.NewStatusCmd(ctx)
	cmd.SetVal("OK")
	return cmd
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	m.mu.RLock()
	defer m.mu.RUnlock()

	data, ok := m.Data[key]
	if !ok {
		return redis.NewStringResult("", redis.Nil)
	}
	return redis.NewStringResult(string(data), redis.Nil)
}

func (m *MockRedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	m.mu.Lock()
	defer m.mu.Unlock()

	var count int64
	for _, key := range keys {
		if _, ok := m.Data[key]; ok {
			delete(m.Data, key)
			count++
		}
	}
	return redis.NewIntResult(count, redis.Nil)
}
func (m *MockRedisClient) Ping(ctx context.Context) *redis.StatusCmd {
	cmd := redis.NewStatusCmd(ctx)
	cmd.SetVal("PONG")
	return cmd
}

func (m *MockRedisClient) Close() error {
	return nil
}
