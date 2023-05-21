package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Login function validate credentials of user and redirect request
func Login(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		return c.Render(http.StatusOK, "login", nil)
	}
	//more config to method post
	return c.Redirect(http.StatusFound, "/app")
}

func App(c echo.Context) error {
	return c.String(200, "application")
}
