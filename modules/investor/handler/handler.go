package handler

import (
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
	param := utilities.Param{}
	param.Limit = 10
	param.Page = 1

	err = (&echo.DefaultBinder{}).BindQueryParams(e, &param)
	if err != nil {
		log.Printf("[ERROR] Handler GetAll BindQueryParam: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.ErrorResponse(err.Error()))
	}

	investors, err := h.usecase.GetAll(ctx, &param)
	if err != nil {
		log.Printf("[ERROR] Handler GetAll: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.ErrorResponse("Something went wrong"))
	}
	meta := utilities.BuildMeta(param, len(investors))
	return e.JSON(http.StatusOK, utilities.SuccessResponse(meta, investors))
}

func (h *Handler) GetByID(e echo.Context) (err error) {
	ctx := e.Request().Context()
	id := e.Param("id")
	investorId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("[ERROR] Handler GetByID ParseInt: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.ErrorResponse(err.Error()))
	}
	investor, err := h.usecase.GetByID(ctx, investorId)
	if err != nil {
		if err == sql.ErrNoRows {
			return e.JSON(http.StatusNotFound, utilities.ErrorResponse(err.Error()))
		}
		log.Printf("[ERROR] Handler GetByID: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.ErrorResponse("Something went wrong"))
	}
	return e.JSON(http.StatusOK, utilities.SuccessResponse(nil, investor))
}
