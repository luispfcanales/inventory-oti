package controller

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/ports"
	"github.com/luispfcanales/inventory-oti/services"
)

// PageLogin function validate credentials of user and redirect request
func PageLogin(c *fiber.Ctx) error {
	return c.Render("page_login", nil)
}

// LoginPost authenticate if exist user and create session
func LoginPost(AuthSrv ports.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		u, err := AuthSrv.AuthUser(username, password)
		if err != nil {
			log.Println("Error login Controller: ", err)
			return c.Render("page_login", nil)
		}

		isess := services.GetInstanceSession(c)
		isess.Set(services.KEY_SESSION_USERNAME, fmt.Sprintf("%s %s", u.FirstName, u.LastName))

		if err := isess.Save(); err != nil {
			log.Println(err)
			return c.SendString("internal server error")
		}

		return c.Redirect("/admin", fiber.StatusSeeOther)
	}
}

// LoginExit destroy session and redirect to /login
func LoginExit(c *fiber.Ctx) error {
	isess := services.GetInstanceSession(c)
	name := isess.Get(services.KEY_SESSION_USERNAME)

	isess.Delete(services.KEY_SESSION_USERNAME)
	if err := isess.Destroy(); err != nil {
		panic(err)
	}

	log.Println(fmt.Sprintf("Exit user -> %v", name))
	return c.Redirect("/login", fiber.StatusSeeOther)
}

// Index return page wellcome application
func Index(c *fiber.Ctx) error {
	return c.Render("index", nil)
}
