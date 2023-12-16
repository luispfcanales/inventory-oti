package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func HdlGetAllNetworks(networkSrv ports.NetworkService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)
		list := networkSrv.ListNetworks()
		return res.SendJSON(list)
	}
}

func HdlGetResumeNetworks(networkSrv ports.NetworkService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)
		list := networkSrv.ListResumeNetworks()
		return res.SendJSON(list)
	}
}
