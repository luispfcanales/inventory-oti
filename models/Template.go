package models

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	EngineTemplate *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.EngineTemplate.ExecuteTemplate(w, "login", nil)
}
