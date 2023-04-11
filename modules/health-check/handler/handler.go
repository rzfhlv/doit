package handler

import (
	"context"
	"doit/modules/health-check/usecase"
	"doit/utilities"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IHandler interface {
	HealthCheck(e echo.Context) (err error)
}

type Handler struct {
	usecase usecase.IUsecase
}

func NewHandler(usecase usecase.IUsecase) IHandler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) HealthCheck(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()

	err = h.usecase.HealthCheck(ctx)
	if err != nil {
		log.Printf("[ERROR] Handler Health Check: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.SetResponse("error", "Something went wrong", nil, nil))
	}
	return e.JSON(http.StatusOK, utilities.SetResponse("ok", "I'm health", nil, nil))
}
