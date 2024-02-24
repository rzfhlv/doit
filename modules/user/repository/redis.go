package repository

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
)

func (r *Repository) Set(ctx context.Context, key string, value string, ttl time.Duration) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "User Repository Redis Set")
	defer sp.Finish()

	err = r.redis.Set(ctx, key, value, ttl).Err()
	return
}

func (r *Repository) Get(ctx context.Context, key string) (value string, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "User Repository Redis Get")
	defer sp.Finish()

	value, err = r.redis.Get(ctx, key).Result()
	return
}

func (r *Repository) Del(ctx context.Context, key string) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "User Repository Redis Del")
	defer sp.Finish()

	err = r.redis.Del(ctx, key).Err()
	return
}
