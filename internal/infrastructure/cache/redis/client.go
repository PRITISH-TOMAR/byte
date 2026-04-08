package redis

import (
	"context"
	"fmt"
	"log"

	"by_te/internal/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), cfg.RedisTimeout)
	defer cancel()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal("Redis connection error:", err)
	}

	fmt.Println("Connected to Redis")

	return rdb
}