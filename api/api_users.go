package api

import (
	"github.com/labstack/echo/v4"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func GetUserInfoByID(authSrv ports.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := models.NewResponseApi(c)
		return res.NotFoundJSON()
	}
}
