package api

import (
	"fmt"

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

		if user.Password == "" || user.Username == "" {
			return res.NotFoundJSON()
		}

		u, err := authSrv.AuthUser(user.Username, user.Password)
		if err != nil {
			return res.NotFoundJSON()
		}

		if u.Key == "" {
			return res.NotFoundJSON()
		}

		return res.SendJSON(&models.UserResponse{
			ID:          u.Key,
			AccessToken: u.AccessToken,
			Username:    fmt.Sprintf("%s %s", u.FirstName, u.LastName),
		})
	}
}
