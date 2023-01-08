package main

import (
  "os"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

  e.Static("/assets", "resources/assets")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})


  confAppPort := os.Getenv("APP_PORT")
  if confAppPort == "" {
    confAppPort = "3000"
  }

	e.Logger.Fatal(e.Start(":" + confAppPort))
}

