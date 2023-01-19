package investor

import (
	"doit/modules/investor/handler"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler) (e *echo.Group) {
	e = route.Group("/investor")
	e.GET("", h.GetAll)
	e.GET("/:id", h.GetByID)
	return
}
