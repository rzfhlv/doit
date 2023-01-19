package person

import (
	"doit/modules/person/handler"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler) (e *echo.Group) {
	e = route.Group("/person")
	e.GET("", h.GetAll)
	return
}
