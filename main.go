package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"go-limiter/limiter"
	"net/http"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:63790",
	})

	rateLimiter := limiter.NewRateLimiterService(rdb,
		limiter.WithIdentifierLimiterCustom(2, time.Minute),
		limiter.WithGlobalLimiterCustom(1, time.Minute))

	e := echo.New()
	e.Use(limiter.RateLimitMiddleware(rateLimiter))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Start(":8080")
}
