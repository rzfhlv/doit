package redis

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
	redisErr    error
)

type Redis struct {
	client *redis.Client
}

func NewRedis() (*Redis, error) {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})

		err := redisClient.Ping(context.Background()).Err()
		if err != nil {
			redisErr = err
		}
	})

	if redisErr != nil {
		return nil, redisErr
	}

	return &Redis{
		client: redisClient,
	}, nil
}

func (r *Redis) GetClient() *redis.Client {
	return r.client
}

func (r *Redis) Close() {
	if r.client != nil {
		r.client.Close()
	}
}
