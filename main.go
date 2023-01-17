package main

import (
  //"log"
	"net/http"
  //"database/sql"

	"github.com/labstack/echo/v4"

	"antrianmu-golang/web/config"
	"antrianmu-golang/web/common"
  eventController "antrianmu-golang/web/controllers/event"
)

func main() {
	e := echo.New()

	e.Renderer = common.NewTemplateRenderer("resources/views/*.html")
	e.Static("/assets", "resources/assets")

	//e.GET("/", func(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	//})

  config.Load()
  config.InitDb()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"APP_URL": config.ConfAppUrl,
		})
	}).Name = "homepage"

  e.POST("/event", eventController.Create).Name = "event.create"

	e.Logger.Fatal(e.Start(":" + config.ConfAppPort))
}
