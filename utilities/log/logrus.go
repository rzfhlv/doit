package log

import (
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var (
	LAYOUT = "2006-01-02 15:04:05"
)

func Log(c echo.Context) *log.Entry {
	now := time.Now().Format(LAYOUT)
	if c == nil {
		return log.WithFields(log.Fields{
			"at": now,
		})
	}

	return log.WithFields(log.Fields{
		"at":     now,
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
		"id":     c.Response().Header().Get(echo.HeaderXRequestID),
	})
}
