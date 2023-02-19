package handler

import (
	"context"
	"database/sql"
	"doit/modules/user/model"
	"doit/modules/user/usecase"
	"doit/utilities"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type IHandler interface {
	Register(e echo.Context) (err error)
	Login(e echo.Context) (err error)
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
	ctx := e.Request().WithContext(context.Background()).Context()

	user := model.User{}
	err = e.Bind(&user)
	if err != nil {
		log.Printf("[ERROR] Handler Register Binding: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.SetResponse("error", err.Error(), nil, nil))
	}

	err = e.Validate(user)
	if err != nil {
		log.Printf("[ERROR] Handler Register Validation: %v", err.(validator.ValidationErrors))
		return e.JSON(http.StatusBadRequest, utilities.SetResponse("error", err.Error(), nil, nil))
	}

	result, err := h.usecase.Register(ctx, user)
	if err != nil {
		log.Printf("[ERROR] Handler Register Usecase: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.SetResponse("error", "Something went wrong", nil, nil))
	}
	return e.JSON(http.StatusOK, utilities.SetResponse("ok", "success", nil, result))
}

func (h *Handler) Login(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()

	login := model.Login{}
	err = e.Bind(&login)
	if err != nil {
		log.Printf("[ERROR] Handler Register Binding: %v", err.Error())
		return e.JSON(http.StatusUnprocessableEntity, utilities.SetResponse("error", err.Error(), nil, nil))
	}

	err = e.Validate(login)
	if err != nil {
		log.Printf("[ERROR] Handler Register Validation: %v", err.(validator.ValidationErrors))
		return e.JSON(http.StatusBadRequest, utilities.SetResponse("error", err.Error(), nil, nil))
	}

	result, err := h.usecase.Login(ctx, login)
	if err != nil {
		if err == sql.ErrNoRows || err == bcrypt.ErrMismatchedHashAndPassword {
			log.Printf("[ERROR] Handler Register Usecase: %v", err.Error())
			return e.JSON(http.StatusBadRequest, utilities.SetResponse("error", "Wrong Username or Password", nil, nil))
		}
		log.Printf("[ERROR] Handler Register Usecase: %v", err.Error())
		return e.JSON(http.StatusInternalServerError, utilities.SetResponse("error", "Something went wrong", nil, nil))
	}
	return e.JSON(http.StatusOK, utilities.SetResponse("ok", "success", nil, result))
}
