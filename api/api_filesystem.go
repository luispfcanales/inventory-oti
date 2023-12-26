package api

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
)

var (
	directoryServiceSheet string = "/mnt/s3/ServicesSheet"
	directoryHelper       string = "/mnt/s3/Helper-img"
)

func HdlGetFileS3(c *fiber.Ctx) error {

	keyfile := c.Params("keyfile")
	fullpath := fmt.Sprintf(
		"%s/%s",
		directoryServiceSheet,
		keyfile,
	)

	return c.SendFile(fullpath)
}

func HdlPostFileS3(c *fiber.Ctx) error {
	pathS3 := "/mnt/s3"

	log.Println("se inicia proceso de subida de documento")

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	files := form.File["file"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer src.Close()

		fileType, err := getMimeType(src)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		dst, err := os.Create(filepath.Join(pathS3, file.Filename))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		fmt.Printf("Nombre del archivo: %s, Tipo de archivo: %s\n", file.Filename, fileType)
	}

	return c.SendStatus(fiber.StatusOK)
}
func getMimeType(file io.Reader) (string, error) {
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return "", err
	}
	return mime.String(), nil
}
