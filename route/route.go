package route

import (
	"doit/service"

	"github.com/labstack/echo/v4"
)

func ListRoute(svc *service.Service) (e *echo.Echo) {
	e = echo.New()

	v1 := e.Group("/v1")
	v1.GET("/investor", svc.InvestorHandler.GetAll)
	v1.GET("/investor/:id", svc.InvestorHandler.GetByID)

	v1.GET("/person", svc.PersonHandler.GetAll)

	return
}
