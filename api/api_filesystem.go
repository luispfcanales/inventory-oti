package api

import "github.com/gofiber/fiber/v2"

func HdlGetFile(c *fiber.Ctx) error {
	return c.SendString("hola")
}
