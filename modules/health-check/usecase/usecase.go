package usecase

import (
	"context"
	"fmt"

	"github.com/rzfhlv/doit/modules/health-check/repository"
	logrus "github.com/rzfhlv/doit/utilities/log"
)

type IUsecase interface {
	HealthCheck(ctx context.Context) (err error)
}

type Usecase struct {
	repo repository.IRepository
}

func NewUsecase(repo repository.IRepository) IUsecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) HealthCheck(ctx context.Context) (err error) {
	err = u.repo.Ping(ctx)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Health Check Usecase Postgres Ping, %v", err.Error()))
		return
	}

	err = u.repo.MongoPing(ctx)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Health Check Usecase Mongo Ping, %v", err.Error()))
		return
	}

	err = u.repo.RedisPing(ctx)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Health Check Usecase Redis Ping, %v", err.Error()))
		return
	}
	return
}
