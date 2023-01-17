package config

import (
  "os"
  "database/sql"
)

var (
	ConfAppPort = os.Getenv("APP_PORT")
	ConfAppUrl = os.Getenv("APP_URL")
	ConfDbUrl = os.Getenv("DB_URL")
  DbConn *sql.DB
)

func Load() {
	if ConfAppPort == "" {
		ConfAppPort = "3000"
	}

	if ConfAppUrl == "" {
		ConfAppUrl = "http://localhost:" + ConfAppPort
	}

	if ConfDbUrl == "" {
	  ConfDbUrl = "database.sqlite"
	}
}
