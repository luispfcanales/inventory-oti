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

	s := NewServer(app)
	s.Run()
	//e.Logger.Fatal(s.Run())
}
