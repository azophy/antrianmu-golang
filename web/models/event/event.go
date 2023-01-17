package event

import (
  "database/sql"
  "errors"

  "github.com/mattn/go-sqlite3"
)


type Event struct {
  ID   int64
  Title string
  Description  string
  //Slug string
}

var (
  tablename = "queue_event"
  ErrDuplicate    = errors.New("record already exists")
  //ErrNotExists    = errors.New("row not exists")
  //ErrUpdateFailed = errors.New("update failed")
  //ErrDeleteFailed = errors.New("delete failed")
)


type EventRepository struct {
  db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
  return &EventRepository{
    db: db,
  }
}


func (r *EventRepository) Migrate() error {
  query := `
    CREATE TABLE IF NOT EXISTS queue_event(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL
  );
  `

  _, err := r.db.Exec(query)
  return err
}

func (r *EventRepository) Create(event Event) (*Event, error) {
  res, err := r.db.Exec("INSERT INTO queue_event(title, description) values(?,?)", event.Title, event.Description)
  if err != nil {
    var sqliteErr sqlite3.Error
    if errors.As(err, &sqliteErr) {
      if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
        return nil, ErrDuplicate
      }
    }
    return nil, err
  }

  id, err := res.LastInsertId()
  if err != nil {
    return nil, err
  }
  event.ID = id

  return &event, nil
}

