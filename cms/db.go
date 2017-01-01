package cms

import (
  "database/sql"

  _ "github.com/lib/pq"
)

type PgStore struct {
  DB *sql.DB
}

func newDB() *PgStore {
  db, err:= sql.Open("postgres", "user=nanoo dbname=nanootest sslmode=disable")
  if err != nil {
    panic(err)
  }

  return &PgStore{
    DB: db,
  }
}

func GetPage(id string) (*Page, error) {
  var p Page
  err := store.DB.QuertRow("SELECT * FROM pages WHERE id = $1", id).Scan(&p.ID, &p.Title, &p.Content)
  return &p, err
}

func GetPages() ([]*Page, error) {
  rows, err :=store.DB.Query("SELECT * FROM pages")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  pages := []*Page{}
  for rows.Next() {
    var p Page
    err = rows.Scan(&p.ID, &p.Title, &p.Content)
    if err != nil {
      return nil, err
    }
    pages = append(pages, &p)
  }
  return pages, nil
}

func CreatePage(p *Page) (int, error) {
  var id int
  err := store.DB.QueryRow("INSERT INTO pages(title, content) VALUES($1, $2) RETUENING id", p.Title, p.Content).Scan(&id)
  return id, err
}
