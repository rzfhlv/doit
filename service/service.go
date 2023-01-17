package service

import (
	"doit/config"
	"doit/modules/investor/handler"
	"doit/modules/investor/repository"
	"doit/modules/investor/usecase"
)

type Service struct {
	Investor        usecase.IUsecase
	InvestorHandler handler.IHandler
}

func NewService(cfg *config.Config) *Service {
	investorRepo := repository.NewRepository(cfg.Postgres, cfg.Mongo)
	investorUsecase := usecase.NewUsecase(investorRepo)
	investorHandler := handler.NewHandler(investorUsecase)

	return &Service{
		Investor:        investorUsecase,
		InvestorHandler: investorHandler,
	}
}
