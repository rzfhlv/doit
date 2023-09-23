package usecase

import (
	"context"

	"github.com/rzfhlv/doit/modules/health-check/repository"
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
		return
	}

	err = u.repo.MongoPing(ctx)
	if err != nil {
		return
	}

	err = u.repo.RedisPing(ctx)
	if err != nil {
		return
	}
	return
}
