package repository

import (
	"context"

	"github.com/rzfhlv/doit/modules/user/model"
)

func (r *Repository) Register(ctx context.Context, user model.User) (result model.User, err error) {
	err = r.db.Get(&result, RegisterQuery, user.Name, user.Email, user.Username, user.Password, user.CreatedAt)
	return
}

func (r *Repository) Login(ctx context.Context, login model.Login) (result model.User, err error) {
	err = r.db.Get(&result, LoginQuery, login.Username)
	return
}
