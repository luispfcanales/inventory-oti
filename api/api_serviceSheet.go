package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/ports"
)

func PreviewPDF(fileSrv ports.FileService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/pdf")
		return c.Send(fileSrv.PreviewRenderPDF())
	}
}
