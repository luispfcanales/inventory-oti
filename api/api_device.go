package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func HdlGetAllDevice(networkSrv ports.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)
		list := networkSrv.ListAllDevice()
		return res.SendJSON(list)
	}
}
