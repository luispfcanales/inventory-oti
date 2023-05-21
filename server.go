package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

type server struct {
	engine *echo.Echo
}

// NewServer return instance server with engine
func NewServer(e *echo.Echo) *server {
	return &server{
		engine: e,
	}
}

// Run method up server
func (s *server) Run() error {
	ConfigRoutes(s.engine)
	return s.engine.Start(s.getPort())
}

func (s *server) getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":3000"
	}
	return ":" + port
}
