package middle

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/luispfcanales/inventory-oti/ports"
)

func CheckCookie(AuthSrv ports.AuthService, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cook, err := c.Cookie("Authorization")
		if err != nil {
			log.Println(" no pass =>", err)
			return c.Render(200, "login", nil)
		}
		if !AuthSrv.ValidateTokenCookie(cook.Value) {
			return c.Render(200, "login", nil)
		}
		return next(c)
	}
}

func CheckHeaderToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			tokenString = c.QueryParam("token")
		}

		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
		return next(c)
	}
}
