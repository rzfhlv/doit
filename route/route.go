package route

import (
	healthCheck "doit/modules/health-check"
	"doit/modules/investor"
	"doit/modules/person"
	"doit/modules/user"
	"doit/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func ListRoute(svc *service.Service) (e *echo.Echo) {
	e = echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	route := e.Group("/v1")

	user.Mount(route, svc.UserHandler, svc.AuthMiddleware)
	investor.Mount(route, svc.InvestorHandler, svc.AuthMiddleware)
	person.Mount(route, svc.PersonHandler, svc.AuthMiddleware)
	healthCheck.Mount(route, svc.HelathCheckHandler)

	return
}
