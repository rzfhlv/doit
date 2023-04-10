package auth

import (
	"context"
	"doit/utilities"
	"doit/utilities/jwt"
	"log"
	"net/http"
	"strings"

	"doit/config"

	"github.com/labstack/echo/v4"
)

func AuthBearer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headers := Headers{}
		err := (&echo.DefaultBinder{}).BindHeaders(c, &headers)
		if err != nil {
			log.Println("error middleware")
			return c.JSON(http.StatusUnprocessableEntity, utilities.SetResponse("error", err.Error(), nil, nil))
		}

		split := strings.Split(headers.Authorization, " ")
		if len(split) < 2 {
			log.Printf("[ERROR] Auth Unsupported Token: %v", len(split))
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		if split[0] != "Bearer" {
			log.Printf("[ERROR] Auth Unsupported Token: %v", split[0])
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		if split[1] == "" {
			log.Printf("[ERROR] Auth Empty Token: %v", split[1])
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		claims, err := jwt.ValidateToken(split[1])
		if err != nil {
			log.Printf("[ERROR] Auth Validation Invalid: %v", err.Error())
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		redis, err := config.NewRedis()
		if err != nil {
			log.Printf("[ERROR] Auth Redis Not Connected: %v", err.Error())
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		err = redis.Get(context.Background(), split[1]).Err()
		if err != nil {
			log.Printf("[ERROR] Auth Redis Key Deleted: %v", err.Error())
			return c.JSON(http.StatusUnauthorized, utilities.SetResponse("error", "Unauthorized", nil, nil))
		}

		c.Set("id", claims.ID)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)

		return next(c)
	}
}
