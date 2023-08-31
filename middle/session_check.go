package middle

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/services"
)

// CheckSession verified if exists session
func CheckSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("[ fn CheckSession ]")

		isess := services.GetInstanceSession(c)
		v := isess.Get(services.KEY_SESSION_USERNAME)
		if v == nil {
			return c.Render("page_login", nil)
		}
		return c.Next()
	}
}

// NoCacheMiddleware delete pages caching
func NoCacheMiddleware(c *fiber.Ctx) error {
	c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Set("Pragma", "no-cache")
	c.Set("Expires", "0")
	return c.Next()
}

// RedirectAppIsValidSession redirect page if session is valid
func RedirectAppIsValidSession(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("[ fn RedirectAppIsValidSession ]")

		isess := services.GetInstanceSession(c)
		v := isess.Get(services.KEY_SESSION_USERNAME)
		if v == nil {
			//render login page
			return fn(c)
		}
		return c.Redirect("/admin", fiber.StatusSeeOther)
	}
}
