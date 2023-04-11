package healthcheck

import (
	"doit/modules/health-check/handler"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler) (e *echo.Group) {
	e = route.Group("/health-check")
	e.GET("", h.HealthCheck)
	return
}
