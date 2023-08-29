package routes

import (
	"testing"

	"github.com/rzfhlv/doit/config"
	"github.com/rzfhlv/doit/middleware"
	hc "github.com/rzfhlv/doit/modules/health-check"
	"github.com/rzfhlv/doit/modules/investor"
	"github.com/rzfhlv/doit/modules/person"
	"github.com/rzfhlv/doit/modules/user"
	"github.com/rzfhlv/doit/service"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	cfg := config.Config{
		Postgres: nil,
		Mongo:    nil,
		Redis:    nil,
	}
	service := service.Service{
		Investor:    investor.NewInvestor(&cfg),
		Person:      person.NewPerson(&cfg),
		User:        user.NewUser(&cfg),
		HealthCheck: hc.NewHealthCheck(&cfg),
		Middleware:  middleware.NewMiddleware(&cfg),
	}
	r := ListRoutes(&service)
	assert.NotNil(t, r)
}
