package middle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/services"
)

func CheckToken(c *fiber.Ctx) error {
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
