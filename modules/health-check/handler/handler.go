package handler

import (
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/rzfhlv/doit/modules/health-check/usecase"
	"github.com/rzfhlv/doit/utilities/message"
	"github.com/rzfhlv/doit/utilities/response"

	logrus "github.com/rzfhlv/doit/utilities/log"

	"github.com/labstack/echo/v4"
)

var (
	HEALTHCHEKLOG = "Health Check Handler"
)

type IHandler interface {
	HealthCheck(e echo.Context) (err error)
}

type Handler struct {
	usecase usecase.IUsecase
}

func NewHandler(usecase usecase.IUsecase) IHandler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) HealthCheck(e echo.Context) (err error) {
	ctx := e.Request().Context()
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Health Check Handler HealthCheck")
	defer sp.Finish()

	err = h.usecase.HealthCheck(ctx)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf(HEALTHCHEKLOG+" %v", err.Error()))
		return e.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	return e.JSON(http.StatusOK, response.Set(message.OK, message.HEALTHCHECK, nil, nil))
}
