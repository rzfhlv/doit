package repository

import (
	"context"
	"time"
)

func (r *Repository) Set(ctx context.Context, key string, value string, ttl time.Duration) (err error) {
	err = r.redis.Set(ctx, key, value, ttl).Err()
	return
}

func (r *Repository) Get(ctx context.Context, key string) (value string, err error) {
	value, err = r.redis.Get(ctx, key).Result()
	return
}

func (r *Repository) Del(ctx context.Context, key string) (err error) {
	err = r.redis.Del(ctx, key).Err()
	return
}
