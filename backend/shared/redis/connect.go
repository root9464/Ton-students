package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx context.Context

type Cache struct {
	RedisClient *redis.Client
}

var RedisCache Cache

func Connect(url string) (Cache, error) {
	ctx = context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr: url,
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return Cache{}, err
	}

	RedisCache = Cache{
		RedisClient: redisClient,
	}

	return RedisCache, nil
}
