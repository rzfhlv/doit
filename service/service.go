package service

import (
	"doit/config"
	"doit/modules/investor/repository"
	"doit/modules/investor/usecase"
	"fmt"
	"time"
)

type Service struct {
	Investor usecase.IUsecase
}

func NewService(cfg *config.Config) *Service {
	investorRepo := repository.NewRepository(cfg.Postgres, cfg.Mongo)
	investorUsecase := usecase.NewUsecase(investorRepo)
	start := time.Now()
	duration := time.Since(start)
	fmt.Printf("Done in %v seconds\n", duration.Seconds())

	return &Service{
		Investor: investorUsecase,
	}
}
