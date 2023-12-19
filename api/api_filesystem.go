package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var directoryServiceSheet string = "/home/ubuntu/storage-oti-files"

func HdlFileSystemS3(c *fiber.Ctx) error {

	//keyfile := c.Params("keyfile")
	fullpath := fmt.Sprintf(
		"%s/%s",
		directoryServiceSheet,
		"diagrama_modificado.png",
	)

	return c.SendFile(fullpath)
}
