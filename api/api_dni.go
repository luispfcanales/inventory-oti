package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func GetDataDni(dniSrv ports.ApiService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		dni := c.Params("dni")
		if len(dni) != 8 {
			return models.NewResponseApi(c).NotFoundJSON()
		}

		data, err := dniSrv.GetDataByDni(dni)
		if err != nil {
			return res.NotFoundJSON()
		}
		if !data.Success {
			res.Status = http.StatusNotFound
			return res.NotFoundWithDataJSON(data)
		}

		return res.SendJSON(data)
	}
}
