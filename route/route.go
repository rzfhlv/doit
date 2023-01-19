package route

import (
	"doit/modules/investor"
	"doit/modules/person"
	"doit/service"

	"github.com/labstack/echo/v4"
)

func ListRoute(svc *service.Service) (e *echo.Echo) {
	e = echo.New()

	route := e.Group("/v1")

	investor.Mount(route, svc.InvestorHandler)
	person.Mount(route, svc.PersonHandler)

	return
}
