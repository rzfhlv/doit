package investor

import (
	"doit/config"
	"doit/middleware/auth"
	"doit/modules/investor/handler"
	"doit/modules/investor/repository"
	"doit/modules/investor/usecase"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler, a auth.IAuth) (e *echo.Group) {
	e = route.Group("/investors")
	e.Use(a.Bearer)
	e.GET("", h.GetAll)
	e.GET("/:id", h.GetByID)
	e.POST("/generate", h.Generate)
	e.POST("/migrate", h.Migrate)
	return
}

type Investor struct {
	Handler handler.IHandler
}

func NewInvestor(cfg *config.Config) *Investor {
	Repo := repository.NewRepository(cfg.Postgres, cfg.Mongo)
	Usecase := usecase.NewUsecase(Repo)
	Handler := handler.NewHandler(Usecase)

	return &Investor{
		Handler: Handler,
	}
}
