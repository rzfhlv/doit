package repository

import (
	"context"
	"doit/modules/user/model"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	Register(ctx context.Context, user model.User) (result model.User, err error)
	Login(ctx context.Context, login model.Login) (result model.User, err error)
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Register(ctx context.Context, user model.User) (result model.User, err error) {
	now := time.Now()
	err = r.db.Get(&result, RegisterQuery, user.Name, user.Email, user.Username, user.Password, now)
	if err != nil {
		log.Printf("[ERROR] User Repo Register: %v", err.Error())
	}

	return
}

func (r *Repository) Login(ctx context.Context, login model.Login) (result model.User, err error) {
	err = r.db.Get(&result, LoginQuery, login.Username)
	if err != nil {
		log.Printf("[ERROR] User Repo Register: %v", err.Error())
	}
	return
}
