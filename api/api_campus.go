package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func HdlGetCampus(campusSrv ports.CampusService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		id := c.Params("id")
		list, err := campusSrv.ListCampus(id)
		if err != nil {
			return res.NotFoundJSON()
		}

		res.Data = list
		return res.SendCustomJSON(
			http.StatusOK,
			"Success",
		)
	}
}
func HdlGetAllCampus(campusSrv ports.CampusService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		res.Data = campusSrv.ListAllCampus()
		return res.SendCustomJSON(
			http.StatusOK,
			"Success",
		)
	}
}

func HdlPutCampus(campusSrv ports.CampusService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		camp := models.Campus{}
		err := c.BodyParser(&camp)
		if err != nil {
			return res.BadRequestJSON()
		}

		data, err := campusSrv.PutCampus(camp)
		if err != nil {
			return res.NotFoundWithDataJSON(err.Error())
		}

		return res.SendJSON(data)
	}
}
func HdlPostCampus(campusSrv ports.CampusService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		camp := &models.Campus{}
		err := c.BodyParser(camp)
		if err != nil {
			return res.BadRequestJSON()
		}

		err = campusSrv.SaveCampus(camp)
		if err != nil {
			return err
		}

		res.Data = camp
		return res.CreatedJSON()
	}
}

func HdlDeleteCampus(campusSrv ports.CampusService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		id := c.Params("id")
		err := campusSrv.RemoveCampus(id)
		if err != nil {
			return res.NotFoundJSON()
		}

		return res.SendCustomJSON(
			http.StatusOK,
			"Campus Eliminado",
		)
	}
}
