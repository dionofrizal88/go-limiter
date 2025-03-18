package limiter

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// RateLimiterService is a struct holds redis limiter.
type RateLimiterService struct {
	limiter            *redis_rate.Limiter
	globalLimit        int
	globalDuration     time.Duration
	identifierLimit    int
	identifierDuration time.Duration
}

// NewRateLimiterService is a function uses to initiate redis rate limiter.
func NewRateLimiterService(rdb *redis.Client, options ...Options) *RateLimiterService {
	router := &RateLimiterService{
		limiter:            redis_rate.NewLimiter(rdb),
		globalLimit:        100,
		globalDuration:     5 * time.Minute,
		identifierLimit:    3,
		identifierDuration: 5 * time.Minute,
	}

	for _, opt := range options {
		opt(router)
	}

	return router
}

// AllowRequest is a function uses to set request limit can be allowed.
func (r *RateLimiterService) AllowRequest(ctx context.Context, key string, limit int, duration time.Duration) bool {
	res, err := r.limiter.Allow(ctx, key, redis_rate.Limit{
		Rate:   limit,
		Period: duration,
		Burst:  limit,
	})
	if err != nil {
		return false
	}
	return res.Allowed > 0
}

// RateLimitMiddleware is a function uses to create middleware for echo framework.
func RateLimitMiddleware(rateLimiter *RateLimiterService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			globalKey := "global_rate_limit"
			if !rateLimiter.AllowRequest(ctx, globalKey, rateLimiter.globalLimit, rateLimiter.globalDuration) {
				return c.JSON(http.StatusTooManyRequests, map[string]string{"error": "system is busy"})
			}

			ip := c.RealIP()
			if !rateLimiter.AllowRequest(ctx, ip, rateLimiter.identifierLimit, rateLimiter.identifierDuration) {
				return c.JSON(http.StatusTooManyRequests, map[string]string{"error": "too many requests"})
			}
			return next(c)
		}
	}
}
