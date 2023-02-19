package usecase

import (
	"context"
	"doit/modules/user/model"
	"doit/modules/user/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type IUsecase interface {
	Register(ctx context.Context, user model.User) (result model.User, err error)
	Login(ctx context.Context, login model.Login) (result model.User, err error)
}

type Usecase struct {
	repo repository.IRepository
}

func NewUsecase(repo repository.IRepository) IUsecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Register(ctx context.Context, user model.User) (result model.User, err error) {
	err = user.HashedPassword()
	if err != nil {
		log.Printf("[ERROR] User Usecase Register Hashed Password: %v", err.Error())
		return
	}
	result, err = u.repo.Register(ctx, user)
	if err != nil {
		log.Printf("[ERROR] User Usecase Register: %v", err.Error())
		return
	}
	return
}

func (u *Usecase) Login(ctx context.Context, login model.Login) (result model.User, err error) {
	result, err = u.repo.Login(ctx, login)
	if err != nil {
		log.Printf("[ERROR] User Usecase Login: %v", err.Error())
		return
	}
	err = result.VerifyPassword(login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Printf("[ERROR] User Usecase Login Verify Password: %v", err.Error())
		return
	}
	return
}
