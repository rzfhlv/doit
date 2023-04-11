package person

import (
	"doit/middleware/auth"
	"doit/modules/person/handler"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler, am auth.IAuthMiddleware) (e *echo.Group) {
	e = route.Group("/persons")
	e.Use(am.AuthBearer)
	e.GET("", h.GetAll)
	return
}
