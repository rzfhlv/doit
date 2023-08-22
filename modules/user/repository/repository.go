package repository

import (
	"context"
	"time"

	"github.com/rzfhlv/doit/modules/user/model"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type IRepository interface {
	Register(ctx context.Context, user model.User) (result model.User, err error)
	Login(ctx context.Context, login model.Login) (result model.User, err error)
	Set(ctx context.Context, key string, value string, ttl time.Duration) (err error)
	Get(ctx context.Context, key string) (value string, err error)
	Del(ctx context.Context, key string) (err error)
}

type Repository struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewRepository(db *sqlx.DB, redis *redis.Client) IRepository {
	return &Repository{
		db:    db,
		redis: redis,
	}
}
