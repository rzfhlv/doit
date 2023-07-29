package service

import (
	"github.com/rzfhlv/doit/config"
	"github.com/rzfhlv/doit/middleware"
	hc "github.com/rzfhlv/doit/modules/health-check"
	"github.com/rzfhlv/doit/modules/investor"
	"github.com/rzfhlv/doit/modules/person"
	"github.com/rzfhlv/doit/modules/user"
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
