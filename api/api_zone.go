package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func HdlGetZone(zoneSrv ports.ZoneService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		id := c.Params("id")
		list, err := zoneSrv.ListZone(id)
		if err != nil {
			return res.NotFoundJSON()
		}

		return res.SendJSON(list)
	}
}

func HdlGetAllZone(zoneSrv ports.ZoneService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		list := zoneSrv.ListAllZone()
		if len(list) == 0 {
			return res.NotFoundJSON()
		}
		res.Data = list
		return res.SendCustomJSON(
			http.StatusOK,
			"Success",
		)
	}
}
