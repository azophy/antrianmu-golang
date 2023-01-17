package event

import (
  "fmt"
  "log"
  "net/http"

	"github.com/labstack/echo/v4"

  "antrianmu-golang/web/config"
  eventModel "antrianmu-golang/web/models/event"
)

var (
  eventRepo *eventModel.EventRepository
)

func init() {
  config.InitDb()
  eventRepo = eventModel.NewEventRepository(config.DbConn)

  if err := eventRepo.Migrate(); err != nil {
    log.Fatal(err)
  }
}

// create event
func Create(c echo.Context) error {
  // User ID from path `users/:id`
  newEvent := eventModel.Event{
    Title: c.FormValue("title"),
    Description: c.FormValue("description"),
  }

  res, _ := eventRepo.Create(newEvent)

  msg := fmt.Sprintf("Created new event with id: %d", res.ID)
  return c.String(http.StatusOK, msg)
}

// show event for visitor

// show event for organizer
