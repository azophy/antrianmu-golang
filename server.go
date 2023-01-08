package main

import (
  "os"
  "io"
	"net/http"
  "html/template"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

  renderer := &TemplateRenderer{
      templates: template.Must(template.ParseGlob("resources/views/*.html")),
  }
  e.Renderer = renderer

  e.Static("/assets", "resources/assets")

	//e.GET("/", func(c echo.Context) error {
		//return c.String(http.StatusOK, "Hello, World!")
	//})

  confAppPort := os.Getenv("APP_PORT")
  if confAppPort == "" {
    confAppPort = "3000"
  }

  confAppUrl := os.Getenv("APP_URL")
  if confAppUrl == "" {
    confAppUrl = "http://localhost:" + confAppPort
  }

  e.GET("/", func(c echo.Context) error {
      return c.Render(http.StatusOK, "index.html", map[string]interface{}{
          "APP_URL": confAppUrl,
      })
  }).Name = "homepage"

	e.Logger.Fatal(e.Start(":" + confAppPort))
}

