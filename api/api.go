package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

func Documentation(c *fiber.Ctx) error {
	return c.SendString("hola")
}

func Login(authSrv ports.AuthService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		res := models.NewResponseApi(c)

		user := &models.UserRequest{}
		if err := c.BodyParser(user); err != nil {
			return res.BadRequestJSON()
		}

		log.Println(user)
		if user.Password == "" || user.Username == "" {
			return res.NotFoundJSON()
		}

		u, err := authSrv.AuthUser(user.Username, user.Password)
		if err != nil {
			log.Println(err)
			return res.NotFoundJSON()
		}

		c.Set("Authorization", "Bearer "+u.AccessToken)
		return res.SendJSON(&models.UserResponse{
			ID:          u.Person.IDPerson,
			Username:    fmt.Sprintf("%s %s", u.FirstName, u.LastName),
			Role:        u.RoleName,
			Staff:       u.Staff,
			Active:      u.Active,
			AccessToken: u.AccessToken,
		})
	}
}
