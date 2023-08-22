package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	logrus "github.com/rzfhlv/doit/utilities/log"
)

func (r *Repository) Set(ctx context.Context, key string, value string, ttl time.Duration) (err error) {
	err = r.redis.Set(ctx, key, value, ttl).Err()
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Repo Set Redis, %v", err.Error()))
	}
	return
}

func (r *Repository) Get(ctx context.Context, key string) (value string, err error) {
	value, err = r.redis.Get(ctx, key).Result()
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Repo Get Redis, %v", err.Error()))
	}

	return
}

func (r *Repository) Del(ctx context.Context, key string) (err error) {
	err = r.redis.Del(ctx, key).Err()
	if err != nil {
		log.Printf("[ERROR] User Repo Del Redis: %v", err.Error())
		logrus.Log(nil).Error(fmt.Sprintf("User Repo Del Redis, %v", err.Error()))
	}
	return
}
