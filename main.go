package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	s := NewServer(e)
	e.Logger.Fatal(s.Run())
}
