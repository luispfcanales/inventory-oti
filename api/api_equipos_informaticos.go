package api

import (
	"github.com/labstack/echo/v4"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func GetComputers(cpuSrv ports.ComputerService) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := models.NewResponseApi(c)
		list := cpuSrv.ListComputers()
		return res.SendJSON(list)
	}
}
