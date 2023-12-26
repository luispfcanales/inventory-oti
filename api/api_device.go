package api

import (
	"log"

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

func HdlPostDevice(deviceSrv ports.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		d := models.Device{}
		err := c.BodyParser(&d)
		if err != nil {
			return res.BadRequestJSON()
		}

		dev, err := deviceSrv.SaveDevice(d)
		if err != nil {
			return res.BadRequestDataJSON(err.Error())
		}
		log.Println(dev)

		return res.CreatedJSON()
	}
}
