package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/luispfcanales/inventory-oti/api"
	"github.com/luispfcanales/inventory-oti/controller"
	"github.com/luispfcanales/inventory-oti/middle"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
	"github.com/luispfcanales/inventory-oti/services"
	"github.com/luispfcanales/inventory-oti/storage/mem"
)

// any service implement ports
var (
	REPOSITORY ports.StorageService
	AUTH_SRV   ports.AuthService
)

// initialized all services
func init() {
	REPOSITORY = mem.NewStorage()
	AUTH_SRV = services.NewAuth(REPOSITORY)
}

// ConfigRoutes setting routes to api and controllers routes
func ConfigRoutes(e *echo.Echo) {
	e.Static("/public", "public")

	RegisterRoutesController(e)
	CreateApiRoutes(e)
}

// RegisterRoutesController set the routes
func RegisterRoutesController(e *echo.Echo) {
	t := &models.Template{
		EngineTemplate: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	//render application form login
	e.GET("/", controller.Login(AUTH_SRV))

	//authentication application render view
	e.POST("/", controller.Login(AUTH_SRV))
	e.GET("/app", middle.CheckCookie(controller.App))
}

// CreateApiRoutes create new routes to /api/anyroutes
func CreateApiRoutes(e *echo.Echo) {
	a := e.Group("/api")
	a.GET("", api.Documentation)
	a.POST("/login", api.Login(AUTH_SRV))
}
