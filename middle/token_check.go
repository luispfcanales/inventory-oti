package middle

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/services"
)

func CheckToken(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	c.Set("Access-Control-Expose-Headers", "Authorization")
	c.Set("Access-Control-Allow-Credentials", "true")
	c.Set("Access-Control-Max-Age", "3600")
	log.Println("si llego con las acbeceras")

	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	token, err := jwt.ParseWithClaims(tokenString[7:], &models.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return services.Keytoken, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token signature",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}
	claims, ok := token.Claims.(*models.JWTCustomClaims)
	if !ok {
	}

	if !claims.Active {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Usuario desactivado",
		})
	}

	return c.Next()
}
