package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engineTemplates := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engineTemplates,
	})

	s := NewServer(app)
	s.Run()
	//e.Logger.Fatal(s.Run())
}
