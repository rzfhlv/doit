package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rzfhlv/doit/modules/person/usecase"
	"github.com/rzfhlv/doit/utilities/message"
	"github.com/rzfhlv/doit/utilities/param"
	"github.com/rzfhlv/doit/utilities/response"

	logrus "github.com/rzfhlv/doit/utilities/log"

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
	param := param.Param{}
	param.Limit = 10
	param.Page = 1

	err = (&echo.DefaultBinder{}).BindQueryParams(e, &param)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Person Handler GetAll BindQueryParam, %v", err.Error()))
		return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	persons, err := h.usecase.GetAll(ctx, &param)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Person Handler GetAll, %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	meta := response.BuildMeta(param, len(persons))
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, meta, persons))
}

func (h *Handler) GetByID(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()
	id := e.Param("id")
	personId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Person Handler GetByID ParseInt, %v", err.Error()))
		return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
	}
	person, err := h.usecase.GetByID(ctx, personId)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Person Handler GetByID, %v", err.Error()))
		if err == mongo.ErrNoDocuments {
			return e.JSON(http.StatusNotFound, response.Set(message.ERROR, message.NOTFOUND, nil, nil))
		}
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, person))
}
