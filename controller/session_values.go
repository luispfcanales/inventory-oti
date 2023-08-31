package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/services"
)

func SessionValue(c *fiber.Ctx) error {
	var username string
	var isValid bool

	id := c.Params("id")
	isess := services.GetInstanceSession(c)
	if isess.ID() != id {
		return c.Redirect("/login")
	}
	value := isess.Get(services.KEY_SESSION_USERNAME)
	if value != nil {
		username = fmt.Sprintf("%v", value)
		isValid = true
	}
	return c.JSON(map[string]interface{}{
		"IsValid":  isValid,
		"Username": username,
	})
}
