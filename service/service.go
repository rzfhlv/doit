package service

import (
	"doit/config"
	"doit/middleware"
	hc "doit/modules/health-check"
	"doit/modules/investor"
	"doit/modules/person"
	"doit/modules/user"
)

type Service struct {
	Investor    *investor.Investor
	Person      *person.Person
	User        *user.User
	HealthCheck *hc.HealthCheck
	Middleware  *middleware.Middleware
}

func NewService(cfg *config.Config) *Service {
	investorModule := investor.NewInvestor(cfg)
	personModule := person.NewPerson(cfg)
	userModule := user.NewUser(cfg)
	healthCheckModule := hc.NewHealthCheck(cfg)

	middleware := middleware.NewMiddleware(cfg.Redis)

	return &Service{
		Investor:    investorModule,
		Person:      personModule,
		User:        userModule,
		HealthCheck: healthCheckModule,
		Middleware:  middleware,
	}
}
