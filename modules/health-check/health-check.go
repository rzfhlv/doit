package healthcheck

import (
	"github.com/rzfhlv/doit/config"
	"github.com/rzfhlv/doit/modules/health-check/handler"
	"github.com/rzfhlv/doit/modules/health-check/repository"
	"github.com/rzfhlv/doit/modules/health-check/usecase"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler) (e *echo.Group) {
	e = route.Group("/health-check")
	e.GET("", h.HealthCheck)
	return
}

type HealthCheck struct {
	Handler handler.IHandler
}

func NewHealthCheck(cfg *config.Config) *HealthCheck {
	Repo := repository.NewRepository(cfg.Postgres, cfg.Mongo, cfg.Redis)
	Usecase := usecase.NewUsecase(Repo)
	Handler := handler.NewHandler(Usecase)

	return &HealthCheck{
		Handler: Handler,
	}
}
