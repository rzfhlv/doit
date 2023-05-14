package user

import (
	"doit/config"
	"doit/middleware/auth"
	"doit/modules/user/handler"
	"doit/modules/user/repository"
	"doit/modules/user/usecase"

	"github.com/labstack/echo/v4"
)

func Mount(route *echo.Group, h handler.IHandler, a auth.IAuth) (e *echo.Group) {
	e = route.Group("/users")
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.POST("/validate", h.Validate, a.Bearer)
	e.POST("/logout", h.Logout, a.Bearer)
	return
}

type User struct {
	Handler handler.IHandler
}

func NewUser(cfg *config.Config) *User {
	Repo := repository.NewRepository(cfg.Postgres, cfg.Redis)
	Usecase := usecase.NewUsecase(Repo)
	Handler := handler.NewHandler(Usecase)

	return &User{
		Handler: Handler,
	}
}
