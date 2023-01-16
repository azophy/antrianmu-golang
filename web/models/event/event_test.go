package event

import (
  "database/sql"
  "os"
  "testing"

  _ "github.com/mattn/go-sqlite3"
)

const fileName = "test_sqlite.db"

func TestSqlBasic(t *testing.T) {
  os.Remove(fileName)

  db, err := sql.Open("sqlite3", fileName)
  if err != nil {
    t.Log("Failed creating sqlite connection", err)
    t.Fail()
  }

  eventRepo := NewSQLiteRepository(db)
  if err := eventRepo.Migrate(); err != nil {
    t.Log("Failed migrating table creation", err)
    t.Fail()
  }

  newEvent := Event{
    Title: "sesuatu",
    Description: "wow",
  }
  if _,err := eventRepo.Create(newEvent); err != nil {
    t.Log("Failed inserting data", err)
    t.Fail()
  }
}

