package event

import {
  eventModel "antrianmu-golang/web/models/event"
}

// create event
func create(c echo.Context) error {
  // User ID from path `users/:id`
  newEvent := eventModel.event{
    "title": c.FormValue("title"),
    "description": c.FormValue("description"),
  }

  res, _ := eventModel.create(newEvent)

  return c.String(http.StatusOK, res.id)
}

// show event for visitor

// show event for organizer
