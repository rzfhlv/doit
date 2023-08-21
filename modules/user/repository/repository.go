package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rzfhlv/doit/modules/user/model"

	logrus "github.com/rzfhlv/doit/utilities/log"

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

func (r *Repository) Register(ctx context.Context, user model.User) (result model.User, err error) {
	err = r.db.Get(&result, RegisterQuery, user.Name, user.Email, user.Username, user.Password, user.CreatedAt)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Repo Register, %v", err.Error()))
	}

	return
}

func (r *Repository) Login(ctx context.Context, login model.Login) (result model.User, err error) {
	err = r.db.Get(&result, LoginQuery, login.Username)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Repo Login, %v", err.Error()))
	}
	return
}

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
