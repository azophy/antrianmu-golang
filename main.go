package main

import (
	"net/http"
	"os"

	"antrianmu-golang/web/common"
	eventController "antrianmu-golang/web/controllers/event"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Renderer = common.NewTemplateRenderer("resources/views/*.html")
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

	confDbUrl := "database.sqlite"
	db, err := sql.Open("sqlite3", confDbUrl)
	if err != nil {
		log.Fatal(err)
	}

	AppConf := map[string]interface{}{
		"db": db,
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"APP_URL": confAppUrl,
		})
	}).Name = "homepage"

	e.POST("/event", eventController.create)

	e.Logger.Fatal(e.Start(":" + confAppPort))
}
