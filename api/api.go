package api

import (
	"github.com/labstack/echo/v4"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func Documentation(c echo.Context) error {
	return c.String(200, "welcome")
}

func Login(authSrv ports.AuthService) echo.HandlerFunc {

	return func(c echo.Context) error {
		res := models.NewResponseApi(c)

		user := &models.UserRequest{}
		if err := c.Bind(user); err != nil {
			return res.BadRequestJSON()
		}

		if user.Password == "" || user.Username == "" {
			return res.NotFoundJSON()
		}

		u, err := authSrv.AuthUser(user.Username, user.Password)
		if err != nil {
			return res.NotFoundJSON()
		}

		if u.Key == "" {
			return res.NotFoundJSON()
		}

		return res.SendJSON(&models.UserResponse{
			ID:          u.Key,
			AccessToken: u.AccessToken,
		})
	}
}
