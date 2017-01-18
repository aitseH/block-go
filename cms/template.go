package cms

import (
	"html/template"
	"time"
)

var Tmpl = template.Must(template.ParseGlob("./templates/*"))

type Page struct {
	ID      int
	Title   string
	Content string
	Posts   []*Post
}

type Post struct {
	ID            int
	Title         string
	Content       string
	DatePublished time.Time
	Comments      []*Comment
}

type Comment struct {
	Id            int
	Author        string
	Comment       string
	DatePublished time.Time
}
