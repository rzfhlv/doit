package investor

import (
	"doit/middleware/auth"
	"doit/modules/investor/handler"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler) (e *echo.Group) {
	e = route.Group("/investors")
	e.Use(auth.AuthBearer)
	e.GET("", h.GetAll)
	e.GET("/:id", h.GetByID)
	return
}
