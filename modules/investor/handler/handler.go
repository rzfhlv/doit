package handler

import (
	"context"
	"database/sql"
	"doit/modules/investor/usecase"
	"doit/utilities"
	"log"
	"net/http"
	"strconv"

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
	ctx := e.Request().WithContext(context.Background()).Context()
	param := utilities.Param{}
	param.Limit = 10
	param.Page = 1

	err = (&echo.DefaultBinder{}).BindQueryParams(e, &param)
	if err != nil {
		log.Printf("[ERROR] Handler GetAll BindQueryParam: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.SetResponse("error", err.Error(), nil, nil))
	}

	investors, err := h.usecase.GetAll(ctx, &param)
	if err != nil {
		log.Printf("[ERROR] Handler GetAll: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.SetResponse("error", "Something went wrong", nil, nil))
	}
	meta := utilities.BuildMeta(param, len(investors))
	return e.JSON(http.StatusOK, utilities.SetResponse("ok", "success", meta, investors))
}

func (h *Handler) GetByID(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()
	id := e.Param("id")
	investorId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("[ERROR] Handler GetByID ParseInt: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.SetResponse("error", err.Error(), nil, nil))
	}
	investor, err := h.usecase.GetByID(ctx, investorId)
	if err != nil {
		log.Printf("[ERROR] Handler GetByID: %v", err.Error())
		if err == sql.ErrNoRows {
			return e.JSON(http.StatusNotFound, utilities.SetResponse("error", "Not found", nil, nil))
		}
		return e.JSON(http.StatusInternalServerError, utilities.SetResponse("error", "Something went wrong", nil, nil))
	}
	return e.JSON(http.StatusOK, utilities.SetResponse("ok", "success", nil, investor))
}

func (h *Handler) Generate(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()

	err = h.usecase.Generate(ctx)
	if err != nil {
		log.Printf("[ERROR] Handler Generate: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.SetResponse("error", "Something went wrong", nil, nil))
	}
	return e.JSON(http.StatusOK, utilities.SetResponse("ok", "success", nil, nil))
}

func (h *Handler) Migrate(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()

	err = h.usecase.MigrateInvestors(ctx)
	if err != nil {
		log.Printf("[ERROR] Handler Migrate: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.SetResponse("error", "Something went wrong", nil, nil))
	}
	return e.JSON(http.StatusOK, utilities.SetResponse("ok", "success", nil, nil))
}
