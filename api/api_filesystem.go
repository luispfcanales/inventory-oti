package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
)

var directory []string = []string{
	"/mnt/s3/ServicesSheet",
	"/mnt/s3/Helper-img",
	"/mnt/s3/ticket-img",
}

// HdlGetFileS3 return file of S3
//
//	Directory index:
//	0 -> ServicesSheet
//	1 -> Helper-img
//	2 -> ticket-img
func HdlGetFileS3(indexDirectory uint) fiber.Handler {
	return func(c *fiber.Ctx) error {

		keyfile := c.Params("keyfile")
		fullpath := fmt.Sprintf(
			"%s/%s",
			directory[indexDirectory],
			keyfile,
		)

		return c.SendFile(fullpath)
	}
}

func HdlPostFileS3(indexDir uint) fiber.Handler {
	return func(c *fiber.Ctx) error {

		log.Println("se inicia proceso de subida de documento")

		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		files := form.File["load_file"]

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

			dst, err := os.Create(filepath.Join(directory[indexDir], file.Filename))
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
}
func getMimeType(file io.Reader) (string, error) {
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return "", err
	}
	return mime.String(), nil
}

// HdlReturnArgsToFirmPDF return args base64 to send firm refirma invoker
func HdlReturnArgsToFirmPDF(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	res := models.NewResponseApi(c)

	opts := models.OptionsSignature{}
	err := c.BodyParser(&opts)
	if err != nil {
		return res.BadRequestJSON()
	}

	reniec := make(map[string]string)

	log.Println(opts)
	reniec["app"] = "pdf"
	reniec["clientId"] = "ZIzAvpCQernywPNktelaHQH0yi0"
	reniec["clientSecret"] = "B6jWcQmOjJkD94A-EgTl"
	reniec["idFile"] = "load_file"
	reniec["type"] = "W"
	reniec["protocol"] = "T"                                                               //https: S - http: T
	reniec["fileDownloadUrl"] = "http://18.219.214.89/file/reniec/" + opts.FileID + ".pdf" //endpoint
	reniec["fileDownloadLogoUrl"] = ""                                                     //logo
	reniec["fileDownloadStampUrl"] = "http://18.219.214.89/public/logofirma.png"           //stamp reniec logo - optional
	reniec["fileUploadUrl"] = "http://18.219.214.89/file/sheet/upload"                     //route to upload file and save
	reniec["contentFile"] = opts.FileID + ".pdf"                                           //real name document - json struct
	reniec["reason"] = opts.Reason                                                         //json struct
	reniec["pageNumber"] = opts.PageNumber                                                 //json struct
	reniec["posx"] = opts.Pox                                                              //json sctruct
	reniec["posy"] = opts.Poy                                                              //json sctruct
	reniec["isSignatureVisible"] = "true"
	reniec["stampAppearanceId"] = opts.StampAppearanceID //json struct
	reniec["fontSize"] = "7"
	reniec["dcfilter"] = ".*FIR.*|.*FAU.*"
	reniec["outputFile"] = opts.FileID + "[R].pdf" //json struct name file
	reniec["maxFileSize"] = "41943040"             //40Mb
	reniec["timestamp"] = "false"

	b, err := json.Marshal(reniec)
	if err != nil {
		return c.JSON(fiber.Map{
			"err": err,
		})
	}

	ed := base64.StdEncoding.EncodeToString(b)
	return c.JSON(fiber.Map{
		"args": ed,
	})
}

func HdlGetFileReniecS3(indexDirectory uint) fiber.Handler {
	return func(c *fiber.Ctx) error {

		keyfile := c.Params("keyfile")
		fullpath := fmt.Sprintf(
			"%s/%s",
			directory[indexDirectory],
			keyfile,
		)

		fs, err := os.Open(fullpath)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString("Error interno del servidor")
		}
		defer fs.Close()

		c.Set("Content-Type", "application/octet-stream")
		c.Set("Content-Disposition", "attachment; filename="+keyfile)

		if err := c.SendStream(fs); err != nil {
			log.Println(err)
			return c.Status(500).SendString("Error interno del servidor")
		}

		return c.SendStatus(200)
	}
}
