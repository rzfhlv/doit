package repository

import (
	"context"
	"fmt"

	"github.com/rzfhlv/doit/modules/user/model"
	logrus "github.com/rzfhlv/doit/utilities/log"
)

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
