package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func GetUserInfoByID(authSrv ports.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)
		return res.NotFoundJSON()
	}
}

func GetAllUsers(userSrv ports.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		list, err := userSrv.ListAllUsers()
		if err != nil {
			log.Println(err)
			return res.NotFoundJSON()
		}

		return res.SendJSON(list)
	}
}

func HdlGetRole(userSrv ports.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		list, err := userSrv.ListRole()
		if err != nil {
			log.Println(err)
			return res.NotFoundJSON()
		}

		return res.SendJSON(list)
	}
}
func HdlGetStaff(userSrv ports.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		list, err := userSrv.ListStaff()
		if err != nil {
			log.Println(err)
			return res.NotFoundJSON()
		}

		return res.SendJSON(list)
	}
}

func CreateUser(userSrv ports.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		u := models.User{}

		err := c.BodyParser(&u)
		if err != nil {
			return res.BadRequestJSON()
		}

		err = userSrv.SaveUser(&u)
		if err != nil {
			log.Println(err)
			return res.BadRequestDataJSON(err.Error())
		}

		return res.CreatedJSON()
	}
}

func UpdateUser(userSrv ports.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		u := models.User{}

		err := c.BodyParser(&u)
		if err != nil {
			return res.BadRequestJSON()
		}
		log.Println(u)

		return res.CreatedJSON()

		//err = userSrv.Save(u)
		//if err != nil {
		//	log.Println(err)
		//	return res.BadRequestDataJSON(err.Error())
		//}

	}
}
