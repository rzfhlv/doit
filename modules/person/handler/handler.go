package handler

import (
	"context"
	"doit/modules/person/usecase"
	"doit/utilities"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type IHandler interface {
	GetAll(e echo.Context) (err error)
	GetByID(e echo.Context) (err error)
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

func (h *Handler) GetByID(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()
	id := e.Param("id")
	personId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("[ERROR] Handler GetByID ParseInt: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.SetResponse("error", err.Error(), nil, nil))
	}
	person, err := h.usecase.GetByID(ctx, personId)
	if err != nil {
		log.Printf("[ERROR] Handler GetByID: %v", err.Error())
		if err == mongo.ErrNoDocuments {
			return e.JSON(http.StatusNotFound, utilities.SetResponse("error", "Not found", nil, nil))
		}
		return e.JSON(http.StatusInternalServerError, utilities.SetResponse("error", "Something went wrong", nil, nil))
	}
	return e.JSON(http.StatusOK, utilities.SetResponse("ok", "success", nil, person))
}
