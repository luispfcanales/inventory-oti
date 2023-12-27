package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	engine *fiber.App
}

// NewServer return instance server with engine
func NewServer(e *fiber.App) *server {
	return &server{
		engine: e,
	}
}

// Run method up server
func (s *server) Run() error {
	log.SetFlags(log.Lshortfile | log.Ldate)
	ConfigRoutes(s.engine)
	return s.engine.Listen(s.getPort())
}

func (s *server) getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":3000"
	}
	return ":" + port
}
