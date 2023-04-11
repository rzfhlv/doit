package usecase

import (
	"context"
	"doit/modules/health-check/repository"
	"log"
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
		log.Printf("[ERROR] Usecase Health Check Postgres Ping: %v", err.Error())
		return
	}

	err = u.repo.MongoPing(ctx)
	if err != nil {
		log.Printf("[ERROR] Usecase Health Check Mongo Ping: %v", err.Error())
		return
	}

	err = u.repo.RedisPing(ctx)
	if err != nil {
		log.Printf("[ERROR] Usecase Health Check Redis Ping: %v", err.Error())
		return
	}
	return
}
