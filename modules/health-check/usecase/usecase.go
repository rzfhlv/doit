package usecase

import (
	"context"

	"github.com/opentracing/opentracing-go"
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
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Health Check Usecase HealthCheck")
	defer sp.Finish()

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
