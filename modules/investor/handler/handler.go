package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rzfhlv/doit/modules/investor/usecase"
	"github.com/rzfhlv/doit/utilities/message"
	"github.com/rzfhlv/doit/utilities/param"
	"github.com/rzfhlv/doit/utilities/response"

	logrus "github.com/rzfhlv/doit/utilities/log"

	"github.com/labstack/echo/v4"
)

type IHandler interface {
	GetAll(e echo.Context) (err error)
	GetByID(e echo.Context) (err error)
	Generate(e echo.Context) (err error)
	Migrate(e echo.Context) (err error)
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
	ctx := e.Request().Context()
	param := param.Param{}
	param.Limit = 10
	param.Page = 1

	err = (&echo.DefaultBinder{}).BindQueryParams(e, &param)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Handler GetAll BindQueryParam, %v", err.Error()))
		return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	investors, err := h.usecase.GetAll(ctx, &param)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Handler GetAll, %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	meta := response.BuildMeta(param, len(investors))
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, meta, investors))
}

func (h *Handler) GetByID(e echo.Context) (err error) {
	ctx := e.Request().Context()
	id := e.Param("id")
	investorId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Handler GetByID ParseInt, %v", err.Error()))
		return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
	}
	investor, err := h.usecase.GetByID(ctx, investorId)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Handler GetByID, %v", err.Error()))
		if err == sql.ErrNoRows {
			return e.JSON(http.StatusNotFound, response.Set(message.ERROR, message.NOTFOUND, nil, nil))
		}
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, investor))
}

func (h *Handler) Generate(e echo.Context) (err error) {
	ctx := e.Request().Context()

	err = h.usecase.Generate(ctx)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Handler Generate, %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, nil))
}

func (h *Handler) Migrate(e echo.Context) (err error) {
	ctx := e.Request().Context()

	err = h.usecase.MigrateInvestors(ctx)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Handler Migrate, %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, nil))
}
