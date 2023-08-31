package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func GetComputers(cpuSrv ports.ComputerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)
		list := cpuSrv.ListComputers()
		return res.SendJSON(list)
	}
}
