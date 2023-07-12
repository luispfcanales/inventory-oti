package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/luispfcanales/inventory-oti/ports"
)

// Login function validate credentials of user and redirect request
func Login(AuthSrv ports.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		method := c.Request().Method

		switch method {
		case http.MethodGet:
			return c.Redirect(http.StatusFound, "/application")

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
			cookie.Name = "Authorization"
			cookie.Value = u.AccessToken
			cookie.Expires = time.Now().Add(30 * time.Minute)
			c.SetCookie(cookie)
			return c.Redirect(http.StatusFound, "/application")
		}

		return nil
	}
}

func ApplicationRender() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "app", nil)
	}
}

// Index return page wellcome application
func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
