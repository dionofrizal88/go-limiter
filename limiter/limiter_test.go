package limiter_test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"go-limiter/limiter"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:63790",
	})

	ctx := context.Background()
	key := "test_key"

	t.Run("test redis limiter with identifier custom value", func(t *testing.T) {
		identifierLimit := 2

		rateLimiter := limiter.NewRateLimiterService(rdb,
			limiter.WithIdentifierLimiterCustom(identifierLimit, time.Minute),
			limiter.WithGlobalLimiterCustom(100, time.Minute))

		err := rdb.FlushDB(context.TODO()).Err()
		assert.Nil(t, err)

		assert.True(t, rateLimiter.AllowRequest(ctx, key, identifierLimit, time.Minute))
		assert.True(t, rateLimiter.AllowRequest(ctx, key, identifierLimit, time.Minute))
		assert.False(t, rateLimiter.AllowRequest(ctx, key, identifierLimit, time.Minute))
	})

	t.Run("test redis limiter with global custom value", func(t *testing.T) {
		globalLimit := 1

		rateLimiter := limiter.NewRateLimiterService(rdb,
			limiter.WithIdentifierLimiterCustom(2, time.Minute),
			limiter.WithGlobalLimiterCustom(globalLimit, time.Minute))

		err := rdb.FlushDB(context.TODO()).Err()
		assert.Nil(t, err)

		assert.True(t, rateLimiter.AllowRequest(ctx, key, globalLimit, time.Minute))
		assert.False(t, rateLimiter.AllowRequest(ctx, key, globalLimit, time.Minute))
	})
}
