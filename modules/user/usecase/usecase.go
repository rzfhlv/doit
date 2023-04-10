package usecase

import (
	"context"
	"doit/modules/user/model"
	"doit/modules/user/repository"
	"doit/utilities/jwt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type IUsecase interface {
	Register(ctx context.Context, user model.User) (result model.JWT, err error)
	Login(ctx context.Context, login model.Login) (result model.JWT, err error)
	Validate(ctx context.Context, validate model.Validate) (result *jwt.JWTClaim, err error)
	Logout(ctx context.Context, token string) (err error)
}

type Usecase struct {
	repo repository.IRepository
}

func NewUsecase(repo repository.IRepository) IUsecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Register(ctx context.Context, user model.User) (result model.JWT, err error) {
	err = user.HashedPassword()
	if err != nil {
		log.Printf("[ERROR] User Usecase Register Hashed Password: %v", err.Error())
		return
	}

	data, err := u.repo.Register(ctx, user)
	if err != nil {
		log.Printf("[ERROR] User Usecase Register: %v", err.Error())
		return
	}

	token, err := jwt.Generate(data.ID, data.Username, data.Email)
	if err != nil {
		log.Printf("[ERROR] User Usecase Register Generate JWT: %v", err.Error())
		return
	}

	err = u.repo.Set(ctx, token, data.ID, time.Duration(1*time.Hour))
	if err != nil {
		log.Printf("[ERROR] User Usecase Register Set Token to Redis: %v", err.Error())
		return
	}

	result.Token = token
	result.Expired = "1 Hour"
	return
}

func (u *Usecase) Login(ctx context.Context, login model.Login) (result model.JWT, err error) {
	data, err := u.repo.Login(ctx, login)
	if err != nil {
		log.Printf("[ERROR] User Usecase Login: %v", err.Error())
		return
	}

	err = data.VerifyPassword(login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Printf("[ERROR] User Usecase Login Verify Password: %v", err.Error())
		return
	}

	token, err := jwt.Generate(data.ID, data.Username, data.Email)
	if err != nil {
		log.Printf("[ERROR] User Usecase Login Generate JWT: %v", err.Error())
		return
	}

	err = u.repo.Set(ctx, token, data.ID, time.Duration(1*time.Hour))
	if err != nil {
		log.Printf("[ERROR] User Usecase Login Set Token to Redis: %v", err.Error())
		return
	}

	result.Token = token
	result.Expired = "1 Hour"
	return
}

func (u *Usecase) Validate(ctx context.Context, validate model.Validate) (result *jwt.JWTClaim, err error) {
	result, err = jwt.ValidateToken(validate.Token)
	if err != nil {
		log.Printf("[ERROR] User Usecase Validate: %v", err.Error())
		return
	}
	return
}

func (u *Usecase) Logout(ctx context.Context, token string) (err error) {
	err = u.repo.Del(ctx, token)
	if err != nil {
		log.Printf("[ERROR] User Usecase Logout: %v", err.Error())
	}
	return
}
