package user

import (
	"doit/middleware/auth"
	"doit/modules/user/handler"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler, am auth.IAuthMiddleware) (e *echo.Group) {
	e = route.Group("/users")
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.POST("/validate", h.Validate, am.AuthBearer)
	e.POST("/logout", h.Logout, am.AuthBearer)
	return
}
