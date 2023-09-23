package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/rzfhlv/doit/modules/user/model"
	"github.com/rzfhlv/doit/modules/user/repository"
	"github.com/rzfhlv/doit/utilities/hasher"
	"github.com/rzfhlv/doit/utilities/jwt"

	"github.com/rzfhlv/doit/utilities/message"

	"golang.org/x/crypto/bcrypt"
)

type IUsecase interface {
	Register(ctx context.Context, user model.User) (result model.JWT, err error)
	Login(ctx context.Context, login model.Login) (result model.JWT, err error)
	Validate(ctx context.Context, validate model.Validate) (result *jwt.JWTClaim, err error)
	Logout(ctx context.Context, token string) (err error)
}

type Usecase struct {
	repo    repository.IRepository
	jwtImpl jwt.JWTInterface
	hasher  hasher.HashPassword
}

func NewUsecase(repo repository.IRepository, jwtImpl jwt.JWTInterface, hasher hasher.HashPassword) IUsecase {
	return &Usecase{
		repo:    repo,
		jwtImpl: jwtImpl,
		hasher:  hasher,
	}
}

func (u *Usecase) Register(ctx context.Context, user model.User) (result model.JWT, err error) {
	hashPassword, err := u.hasher.HashedPassword(user.Password)
	if err != nil {
		return
	}
	user.Password = hashPassword

	checkUser := model.Login{
		Username: user.Username,
		Password: user.Password,
	}
	_, err = u.repo.Login(ctx, checkUser)
	if err == nil {
		err = errors.New(message.ERRUSERNAMEEXIST)
		return
	}

	data, err := u.repo.Register(ctx, user)
	if err != nil {
		return
	}

	token, err := u.jwtImpl.Generate(data.ID, data.Username, data.Email)
	if err != nil {
		return
	}

	err = u.repo.Set(ctx, token, data.Username, time.Duration(1*time.Hour))
	if err != nil {
		return
	}

	result.Token = token
	result.Expired = fmt.Sprintf("%s Hour", os.Getenv("JWT_EXPIRED"))
	return
}

func (u *Usecase) Login(ctx context.Context, login model.Login) (result model.JWT, err error) {
	data, err := u.repo.Login(ctx, login)
	if err != nil {
		return
	}

	err = u.hasher.VerifyPassword(data.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return
	}

	token, err := u.jwtImpl.Generate(data.ID, data.Username, data.Email)
	if err != nil {
		return
	}

	err = u.repo.Set(ctx, token, data.Username, time.Duration(1*time.Hour))
	if err != nil {
		return
	}

	result.Token = token
	result.Expired = fmt.Sprintf("%s Hour", os.Getenv("JWT_EXPIRED"))
	return
}

func (u *Usecase) Validate(ctx context.Context, validate model.Validate) (result *jwt.JWTClaim, err error) {
	result, err = u.jwtImpl.ValidateToken(validate.Token)
	return
}

func (u *Usecase) Logout(ctx context.Context, token string) (err error) {
	err = u.repo.Del(ctx, token)
	return
}
