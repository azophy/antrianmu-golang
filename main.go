package main

import (
  "log"
	"net/http"
  "database/sql"

	"github.com/labstack/echo/v4"
  _ "github.com/mattn/go-sqlite3"

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

	db, err := sql.Open("sqlite3", config.ConfDbUrl)
	if err != nil {
		log.Fatal(err)
	}

  config.DbConn = db

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"APP_URL": config.ConfAppUrl,
		})
	}).Name = "homepage"

  e.POST("/event", eventController.Create).Name = "event.create"

	e.Logger.Fatal(e.Start(":" + config.ConfAppPort))
}
