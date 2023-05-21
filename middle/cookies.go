package middle

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("auth-key")
		if err != nil {
			log.Println(err)
			return c.Redirect(http.StatusFound, "/")
		}
		//more validation of the cookie
		return next(c)
	}
}
