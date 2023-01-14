package main

import (
  "os"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

  e.Renderer = NewTemplateRenderer("resources/views/*.html")
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

