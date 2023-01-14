// adapted from https://echo.labstack.com/guide/templates/
package main

import (
  "io"
  "html/template"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(path string) *TemplateRenderer {
  return &TemplateRenderer{
      templates: template.Must(template.ParseGlob(path)),
  }
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

