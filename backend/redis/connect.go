package redis_connect

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx context.Context

func Connect(url string) (*redis.Client, error) {
	ctx = context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr: url,
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
