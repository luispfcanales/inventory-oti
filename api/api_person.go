package api

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func HdlDeletePerson(personSrv ports.PersonService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		dni := c.Params("dni")
		if len(dni) != 8 {
			return res.NotFoundJSON()
		}

		v, _ := strconv.Atoi(dni)

		if !personSrv.RemovePerson(v) {
			return res.NotFoundJSON()
		}

		return res.SendCustomJSON(http.StatusOK, "Persona eliminada")
	}
}
func HdlPostPerson(personSrv ports.PersonService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		u := models.Person{}
		err := c.BodyParser(&u)
		if err != nil {
			return res.BadRequestJSON()
		}

		_, err = personSrv.SavePerson(u)
		if err != nil {
			return res.SendCustomJSON(
				http.StatusConflict,
				"Dni ya registrado",
			)
		}

		return res.CreatedJSON()
	}
}
func HdlPutPerson(personSrv ports.PersonService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		u := models.Person{}
		err := c.BodyParser(&u)
		if err != nil {
			return res.BadRequestJSON()
		}
		return res.SendJSON(u)
	}
}
func HdlGetPerson(personSrv ports.PersonService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		dni := c.Params("dni")
		if len(dni) != 8 {
			return res.NotFoundJSON()
		}

		v, _ := strconv.Atoi(dni)

		p, err := personSrv.ListOne(v)
		if err != nil {
			return res.NotFoundJSON()
		}

		return res.SendJSON(p)
	}
}
func HdlGetAllPersons(personSrv ports.PersonService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)
		list := personSrv.ListAll()
		return res.SendJSON(list)
	}
}
