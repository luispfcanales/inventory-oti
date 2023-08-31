package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func GetAllConnectionStream(stream ports.StramingComputerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var list []models.StreamEvent
		result := make(chan models.StreamEvent)

		res := models.NewResponseApi(c)

		stream.Receiver(func() {
			stream.ListAllConnections(result)
		})

	loop:
		for {
			select {
			case value, ok := <-result:
				if !ok {
					break loop
				}
				list = append(list, value)
			}
		}

		return res.SendJSON(list)
	}
}
