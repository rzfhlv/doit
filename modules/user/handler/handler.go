package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/rzfhlv/doit/modules/user/model"
	"github.com/rzfhlv/doit/modules/user/usecase"
	"github.com/rzfhlv/doit/utilities/message"
	"github.com/rzfhlv/doit/utilities/response"

	logrus "github.com/rzfhlv/doit/utilities/log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	REGISTERBINDINGLOG    = "User Handler Register Binding"
	REGISTERVALIDATIONLOG = "User Handler Register Validation"
	REGISTERLOG           = "User Handler Register Usecase"

	LOGINBINDINGLOG    = "User Handler Login Binding"
	LOGINVALIDATIONLOG = "User Handler Login Validation"
	LOGINLOG           = "User Handler Login Usecase"

	VALIDATEBINDINGLOG    = "User Handler Validate Binding"
	VALIDATEVALIDATIONLOG = "User Handler Validate Validation"
	VALIDATELOG           = "User Handler Validate Usecase"

	TOKENINVALIDLOG = "User Handler Split Token Invalid"
	LOGOUTLOG       = "User Handler Logout"

	AUTHORIZATION = "Authorization"
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
	ctx := e.Request().Context()
	sp, ctx := opentracing.StartSpanFromContext(ctx, "User Handler Register")
	defer sp.Finish()

	user := model.User{}
	err = e.Bind(&user)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf(REGISTERBINDINGLOG+" %v", err.Error()))
		return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	err = e.Validate(user)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf(REGISTERVALIDATIONLOG+" %v", err.Error()))
		return e.JSON(http.StatusBadRequest, response.Set(message.ERROR, err.Error(), nil, nil))
	}
	user.CreatedAt = time.Now()

	result, err := h.usecase.Register(ctx, user)
	if err != nil {
		if err.Error() == message.ERRUSERNAMEEXIST {
			logrus.Log(nil).Error(fmt.Sprintf(REGISTERLOG+" %v", err.Error()))
			return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, message.USERNAMEEXIST, nil, nil))
		}
		logrus.Log(nil).Error(fmt.Sprintf(REGISTERLOG+" %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, result))
}

func (h *Handler) Login(e echo.Context) (err error) {
	ctx := e.Request().Context()
	sp, ctx := opentracing.StartSpanFromContext(ctx, "User Handler Login")
	defer sp.Finish()

	login := model.Login{}
	err = e.Bind(&login)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf(LOGINBINDINGLOG+" %v", err.Error()))
		return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	err = e.Validate(login)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf(LOGINVALIDATIONLOG+" %v", err.Error()))
		return e.JSON(http.StatusBadRequest, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	result, err := h.usecase.Login(ctx, login)
	if err != nil {
		if err == sql.ErrNoRows || err == bcrypt.ErrMismatchedHashAndPassword {
			logrus.Log(nil).Error(fmt.Sprintf(LOGINLOG+" %v", err.Error()))
			return e.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
		}
		logrus.Log(nil).Error(fmt.Sprintf(LOGINLOG+" %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, result))
}

func (h *Handler) Validate(e echo.Context) (err error) {
	ctx := e.Request().Context()
	sp, ctx := opentracing.StartSpanFromContext(ctx, "User Handler Validate")
	defer sp.Finish()

	validate := model.Validate{}
	err = e.Bind(&validate)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf(VALIDATEBINDINGLOG+" %v", err.Error()))
		return e.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	err = e.Validate(validate)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf(VALIDATEVALIDATIONLOG+" %v", err.(validator.ValidationErrors)))
		return e.JSON(http.StatusBadRequest, response.Set(message.ERROR, err.Error(), nil, nil))
	}

	result, err := h.usecase.Validate(ctx, validate)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf(VALIDATELOG+" %v", err.Error()))
		return e.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.INVALIDTOKEN, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, result))
}

func (h *Handler) Logout(e echo.Context) (err error) {
	ctx := e.Request().Context()
	sp, ctx := opentracing.StartSpanFromContext(ctx, "User Handler Logout")
	defer sp.Finish()

	split := strings.Split(e.Request().Header.Get(AUTHORIZATION), " ")
	if len(split) < 2 {
		logrus.Log(nil).Error(fmt.Sprintf(TOKENINVALIDLOG+" %v", len(split)))
		return e.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
	}

	err = h.usecase.Logout(ctx, split[1])
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf(LOGOUTLOG+" %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.SUCCESS, nil, nil))
}
