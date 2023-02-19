package service

import (
	"doit/config"
	"doit/modules/investor/handler"
	"doit/modules/investor/repository"
	"doit/modules/investor/usecase"
	pHandler "doit/modules/person/handler"
	pRepo "doit/modules/person/repository"
	pUsecase "doit/modules/person/usecase"
	uHandler "doit/modules/user/handler"
	uRepo "doit/modules/user/repository"
	uUsecase "doit/modules/user/usecase"
)

type Service struct {
	Investor        usecase.IUsecase
	InvestorHandler handler.IHandler
	PersonHandler   pHandler.IHandler
	UserHandler     uHandler.IHandler
}

func NewService(cfg *config.Config) *Service {
	investorRepo := repository.NewRepository(cfg.Postgres, cfg.Mongo)
	investorUsecase := usecase.NewUsecase(investorRepo)
	investorHandler := handler.NewHandler(investorUsecase)

	personRepo := pRepo.NewRepository(cfg.Mongo)
	personUsecase := pUsecase.NewUsecase(personRepo)
	personHandler := pHandler.NewHandler(personUsecase)

	userRepo := uRepo.NewRepository(cfg.Postgres)
	userUsecase := uUsecase.NewUsecase(userRepo)
	userHandler := uHandler.NewHandler(userUsecase)

	return &Service{
		Investor:        investorUsecase,
		InvestorHandler: investorHandler,
		PersonHandler:   personHandler,
		UserHandler:     userHandler,
	}
}
