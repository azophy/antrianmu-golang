package config

import (
  "os"
  "log"
  "database/sql"

  _ "github.com/mattn/go-sqlite3"
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

func InitDb() {
	db, err := sql.Open("sqlite3", ConfDbUrl)
	if err != nil {
		log.Fatal(err)
	}

  DbConn = db
}
