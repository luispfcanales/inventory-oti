package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//engineTemplates := html.New("./views", ".html")

	//app := fiber.New(fiber.Config{
	//	Views: engineTemplates,
	//})
	app := fiber.New()
	app.Use(cors.New())
	//app.Use(cors.New(cors.Config{
	//	AllowOrigins:     "*",
	//	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	//	AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
	//	AllowCredentials: true,
	//	ExposeHeaders:    "Authorization",
	//	MaxAge:           3600,
	//}))

	s := NewServer(app)
	s.Run()
	//e.Logger.Fatal(s.Run())
}
