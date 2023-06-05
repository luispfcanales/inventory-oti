package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/luispfcanales/inventory-oti/ports"
)

// Login function validate credentials of user and redirect request
func Login(AuthSrv ports.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		method := c.Request().Method

		switch method {
		case http.MethodGet:
			cook, err := c.Cookie("auth-key")
			if err != nil {
				log.Println(err)
				return c.Render(http.StatusOK, "login", nil)
			}

			if !AuthSrv.ValidateTokenCookie(cook.Value) {
				return c.Render(http.StatusOK, "login", nil)
			}

			return c.Render(http.StatusOK, "app", nil)

		case http.MethodPost: //auth user
			username := c.FormValue("username")
			password := c.FormValue("password")

			u, err := AuthSrv.AuthUser(username, password)
			if err != nil {
				log.Println(err)
				return c.Render(http.StatusOK, "login", nil)
			}
			//send by cookie access token
			cookie := new(http.Cookie)
			cookie.Name = "auth-key"
			cookie.Value = u.AccessToken
			c.SetCookie(cookie)
			return c.Render(http.StatusOK, "app", nil)
		}

		return nil
	}
}

func App(c echo.Context) error {
	return c.String(200, "application")
}
