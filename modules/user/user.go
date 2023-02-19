package user

import (
	"doit/modules/user/handler"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler) (e *echo.Group) {
	e = route.Group("/users")
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	return
}
