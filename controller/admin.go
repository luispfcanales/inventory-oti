package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luispfcanales/inventory-oti/services"
)

func PageAdmin(c *fiber.Ctx) error {
	return c.Render("page_admin_index", services.GetInstanceSession(c).ID())
}
func PageAdminRegisteredComputers(c *fiber.Ctx) error {
	return c.Render("page_admin_registered_computers", services.GetInstanceSession(c).ID())
}
func PageAdminOnlineComputers(c *fiber.Ctx) error {
	return c.Render("page_admin_online_computers", services.GetInstanceSession(c).ID())
}
func PageAdminUserSystem(c *fiber.Ctx) error {
	return c.Render("page_admin_users", services.GetInstanceSession(c).ID())
}
func PageFirmSheet(c *fiber.Ctx) error {
	return c.Render("firmadigital", nil)
}
