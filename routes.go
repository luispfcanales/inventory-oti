package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/luispfcanales/inventory-oti/controller"
	"github.com/luispfcanales/inventory-oti/middle"
	"github.com/luispfcanales/inventory-oti/models"
)

// ConfigRoutes setting routes to api and controllers routes
func ConfigRoutes(e *echo.Echo) {
	e.Static("/static", "assets")

	RegisterRoutesController(e)
	CreateApiRoutes(e)
}

// RegisterRoutesController set the routes
func RegisterRoutesController(e *echo.Echo) {
	t := &models.Template{
		EngineTemplate: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/", controller.Login)
	e.GET("/app", middle.CheckCookie(controller.App))
}

// CreateApiRoutes create new routes to /api/anyroutes
func CreateApiRoutes(e *echo.Echo) {
	a := e.Group("api")
	a.GET("/user", func(c echo.Context) error {
		return c.String(200, "welcome")
	})
}
