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
	e.Use(svc.Middleware.Log.Logrus)
	e.Validator = &CustomValidator{validator: validator.New()}

	route := e.Group("/v1")

	user.Mount(route, svc.User.Handler, svc.Middleware.Auth)
	investor.Mount(route, svc.Investor.Handler, svc.Middleware.Auth)
	person.Mount(route, svc.Person.Handler, svc.Middleware.Auth)
	healthCheck.Mount(route, svc.HealthCheck.Handler)

	return
}
