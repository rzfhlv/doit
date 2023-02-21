package handler

import (
	"context"
	"doit/modules/person/usecase"
	"doit/utilities"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IHandler interface {
	GetAll(e echo.Context) (err error)
}

type Handler struct {
	usecase usecase.IUsecase
}

func NewHandler(usecase usecase.IUsecase) IHandler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) GetAll(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()
	param := utilities.Param{}
	param.Limit = 10
	param.Page = 1

	err = (&echo.DefaultBinder{}).BindQueryParams(e, &param)
	if err != nil {
		log.Printf("[ERROR] Handler GetAll BindQueryParam: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.SetResponse("error", err.Error(), nil, nil))
	}

	persons, err := h.usecase.GetAll(ctx, &param)
	if err != nil {
		log.Printf("[ERROR] Handler GetAll: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.SetResponse("error", "Something went wrong", nil, nil))
	}
	meta := utilities.BuildMeta(param, len(persons))
	return e.JSON(http.StatusOK, utilities.SetResponse("ok", "success", meta, persons))
}
