package log

import (
	uLog "github.com/rzfhlv/doit/utilities/log"

	"github.com/labstack/echo/v4"
)

type ILog interface {
	Logrus(next echo.HandlerFunc) echo.HandlerFunc
}

type Log struct{}

func NewLog() ILog {
	return &Log{}
}

func (l *Log) Logrus(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uLog.Log(c).Info("Incoming Request")
		return next(c)
	}
}
