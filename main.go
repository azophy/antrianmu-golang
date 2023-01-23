package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"antrianmu-golang/web/common"
	"antrianmu-golang/web/config"
	eventController "antrianmu-golang/web/controllers/event"
)

func main() {
	e := echo.New()

	e.Renderer = common.NewTemplateRenderer("resources/views", ".html")
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
