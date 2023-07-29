package person

import (
	"github.com/rzfhlv/doit/config"
	"github.com/rzfhlv/doit/middleware/auth"
	"github.com/rzfhlv/doit/modules/person/handler"
	"github.com/rzfhlv/doit/modules/person/repository"
	"github.com/rzfhlv/doit/modules/person/usecase"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler, a auth.IAuth) (e *echo.Group) {
	e = route.Group("/persons")
	e.Use(a.Bearer)
	e.GET("", h.GetAll)
	e.GET("/:id", h.GetByID)
	return
}

type Person struct {
	Handler handler.IHandler
}

func NewPerson(cfg *config.Config) *Person {
	Repo := repository.NewRepository(cfg.Mongo)
	Usecase := usecase.NewUsecase(Repo)
	Handler := handler.NewHandler(Usecase)

	return &Person{
		Handler: Handler,
	}
}
