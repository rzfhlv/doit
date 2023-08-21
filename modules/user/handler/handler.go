package handler

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/rzfhlv/doit/modules/user/model"
	"github.com/rzfhlv/doit/modules/user/usecase"
	"github.com/rzfhlv/doit/utilities/message"
	"github.com/rzfhlv/doit/utilities/response"

	logrus "github.com/rzfhlv/doit/utilities/log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type IHandler interface {
	Register(e echo.Context) (err error)
	Login(e echo.Context) (err error)
	Validate(e echo.Context) (err error)
	Logout(e echo.Context) (err error)
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
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Register Binding, %v", err.Error()))
		return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	err = e.Validate(user)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Register Validation: %v", err.Error()))
		return e.JSON(http.StatusBadRequest, response.Set(message.ERROR, err.Error(), nil, nil))
	}
	user.CreatedAt = time.Now()

	result, err := h.usecase.Register(ctx, user)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Register Usecase, %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, result))
}

func (h *Handler) Login(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()

	login := model.Login{}
	err = e.Bind(&login)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Login Binding, %v", err.Error()))
		return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	err = e.Validate(login)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Login Validation, %v", err.Error()))
		return e.JSON(http.StatusBadRequest, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	result, err := h.usecase.Login(ctx, login)
	if err != nil {
		if err == sql.ErrNoRows || err == bcrypt.ErrMismatchedHashAndPassword {
			logrus.Log(nil).Error(fmt.Sprintf("User Handler Login Usecase, %v", err.Error()))
			return e.JSON(http.StatusBadRequest, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
		}
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Login Usecase, %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, result))
}

func (h *Handler) Validate(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()

	validate := model.Validate{}
	err = e.Bind(&validate)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Validate Binding, %v", err.Error()))
		return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	err = e.Validate(validate)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Validate Validation, %v", err.(validator.ValidationErrors)))
		return e.JSON(http.StatusBadRequest, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	result, err := h.usecase.Validate(ctx, validate)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Validate Usecase, %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, result))
}

func (h *Handler) Logout(e echo.Context) (err error) {
	ctx := e.Request().WithContext(context.Background()).Context()

	split := strings.Split(e.Request().Header.Get("Authorization"), " ")
	if len(split) < 2 {
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Split Token Invalid, %v", len(split)))
		return e.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
	}

	err = h.usecase.Logout(ctx, split[1])
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("User Handler Logout, %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, nil))
}
