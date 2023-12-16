package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	//engineTemplates := html.New("./views", ".html")

	//app := fiber.New(fiber.Config{
	//	Views: engineTemplates,
	//})
	app := fiber.New()

	s := NewServer(app)
	s.Run()
	//e.Logger.Fatal(s.Run())
}
