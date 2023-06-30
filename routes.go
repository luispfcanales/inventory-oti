package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luispfcanales/inventory-oti/api"
	"github.com/luispfcanales/inventory-oti/controller"
	"github.com/luispfcanales/inventory-oti/middle"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
	"github.com/luispfcanales/inventory-oti/services"
	"github.com/luispfcanales/inventory-oti/storage/clouddeta"
)

// any service implement ports
var (
	REPOSITORY_USER ports.StorageUserService
	REPOSITORY_CPU  ports.StorageComputerService

	AUTH_SRV ports.AuthService
	CPU_SRV  ports.ComputerService
)

// initialized all services
func init() {
	REPOSITORY := clouddeta.NewCloudDetaStorage("e0ytsyfs3et_F3KZDz938AnuKc62WXBdzjt1WnKrNHh8")

	REPOSITORY_USER = REPOSITORY
	REPOSITORY_CPU = REPOSITORY

	AUTH_SRV = services.NewAuth(REPOSITORY_USER)
	CPU_SRV = services.NewComputer(REPOSITORY_CPU)
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
	a.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	a.GET("", api.Documentation)
	a.POST("/login", api.Login(AUTH_SRV))

	a.GET("/computers", middle.CheckHeaderToken(api.GetComputers(CPU_SRV)))
}
