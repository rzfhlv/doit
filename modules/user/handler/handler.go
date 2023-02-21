package handler

import (
	"database/sql"
	"doit/modules/user/model"
	"doit/modules/user/usecase"
	"doit/utilities"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrWrongUsernamePassword = errors.New("wrong username or password")
)

type IHandler interface {
	Register(e echo.Context) (err error)
	Login(e echo.Context) (err error)
	Validate(e echo.Context) (err error)
}

type Handler struct {
	usecase usecase.IUsecase
}

func NewHandler(usecase usecase.IUsecase) IHandler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Register(e echo.Context) (err error) {
	ctx := e.Request().Context()

	user := model.User{}
	err = e.Bind(&user)
	if err != nil {
		log.Printf("[ERROR] Handler Register Binding: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.ErrorResponse(err))
	}

	err = e.Validate(user)
	if err != nil {
		log.Printf("[ERROR] Handler Register Validation: %v", err.(validator.ValidationErrors))
		return e.JSON(http.StatusBadRequest, utilities.ErrorResponse(err))
	}

	result, err := h.usecase.Register(ctx, user)
	if err != nil {
		log.Printf("[ERROR] Handler Register Usecase: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.ErrorResponse(utilities.ErrSomethingWentWrong))
	}
	return e.JSON(http.StatusOK, utilities.SuccessResponse(nil, result))
}

func (h *Handler) Login(e echo.Context) (err error) {
	ctx := e.Request().Context()

	login := model.Login{}
	err = e.Bind(&login)
	if err != nil {
		log.Printf("[ERROR] Handler Register Binding: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.ErrorResponse(err))
	}

	err = e.Validate(login)
	if err != nil {
		log.Printf("[ERROR] Handler Register Validation: %v", err.(validator.ValidationErrors))
		return e.JSON(http.StatusBadRequest, utilities.ErrorResponse(err))
	}

	result, err := h.usecase.Login(ctx, login)
	if err != nil {
		if err == sql.ErrNoRows || err == bcrypt.ErrMismatchedHashAndPassword {
			log.Printf("[ERROR] Handler Register Usecase: %v", err.Error())
			return e.JSON(http.StatusBadRequest, utilities.ErrorResponse(ErrWrongUsernamePassword))
		}
		log.Printf("[ERROR] Handler Register Usecase: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.ErrorResponse(utilities.ErrSomethingWentWrong))
	}
	return e.JSON(http.StatusOK, utilities.SuccessResponse(nil, result))
}

func (h *Handler) Validate(e echo.Context) (err error) {
	ctx := e.Request().Context()

	validate := model.Validate{}
	err = e.Bind(&validate)
	if err != nil {
		log.Printf("[ERROR] Handler Validate Binding: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.ErrorResponse(err))
	}

	err = e.Validate(validate)
	if err != nil {
		log.Printf("[ERROR] Handler Validate Validation: %v", err.(validator.ValidationErrors))
		return e.JSON(http.StatusBadRequest, utilities.ErrorResponse(err))
	}

	result, err := h.usecase.Validate(ctx, validate)
	if err != nil {
		log.Printf("[ERROR] Handler Validate Usecase: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.ErrorResponse(utilities.ErrSomethingWentWrong))
	}
	return e.JSON(http.StatusOK, utilities.SuccessResponse(nil, result))
}
