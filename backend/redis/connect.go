package redis_connect

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/root9464/Ton-students/shared/logger"
)

var ctx context.Context

func Connect(url string, log *logger.Logger) (*redis.Client, error) {
	ctx = context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr: url,
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	log.Info("✅ Redis connection successfully")
	return redisClient, nil
}
