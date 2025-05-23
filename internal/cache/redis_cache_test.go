package cache

import (
	"To_Do/internal/redis"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRedisCache_SetGetDel(t *testing.T) {
	ctx := context.Background()

	type testData struct {
		Name string
		Age  int
	}

	mockClient := NewMockRedis()
	redisCache := NewRedisCache(&redis.RedisClient{Client: mockClient, TTL: 0})

	key := "test:user:1"
	value := testData{
		Name: "Alice",
		Age:  30,
	}

	t.Run("Set/Get/Del success", func(t *testing.T) {
		// set
		err := redisCache.Set(ctx, key, value)
		require.NoError(t, err)
		//Get
		var got testData
		err = redisCache.Get(ctx, key, &got)
		require.NoError(t, err)
		//Del
		err = redisCache.Del(ctx, key)
		require.NoError(t, err)
		// Get после Del
		err = redisCache.Get(ctx, key, &got)
		require.Error(t, err)

	})

	t.Run("Set with invalid JSON", func(t *testing.T) {
		ch := make(chan int)
		err := redisCache.Set(ctx, "bad:key", ch)
		require.NoError(t, err)
	})

	t.Run("Get with invalid JSON", func(t *testing.T) {
		var gat testData
		mockClient.Set(ctx, "corrupted:key", []byte("not-json"), 0)
		err := redisCache.Get(ctx, "corrupted:key", &gat)
		require.NoError(t, err)
	})
}
