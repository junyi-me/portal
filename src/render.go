package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

var tmpl *template.Template

func init() {
    dirs := []string{
        "tmpl/*.html",
        "tmpl/*/*.html",
    }

    files := []string{}
    for _, dir := range dirs {
        ff, err := filepath.Glob(dir)
        if err != nil {
            panic(err)
        }
        files = append(files, ff...)
    }

    t, err := template.ParseFiles(files...)
    if err != nil {
        panic(err)
    }
    tmpl = t
}

type HomePageData struct {
    Hosted   []Link
    External []Link
}
var homeData = HomePageData{hosted, external}

func renderHome(w http.ResponseWriter) {
    err := tmpl.ExecuteTemplate(w, "home.html", homeData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

type Comment struct {
    ID    int
    Email string
    Name  string
    Text  string
    Date  time.Time
    Liked bool
}

type CommentPageData struct {
    LoggedIn bool
    Comments []Comment
}

func renderComments(w http.ResponseWriter) {
    // TODO query
    data := CommentPageData{true, []Comment{
        {1, "sample@example.com", "John Doe", "This is a comment.", time.Now(), false},
        {2, "sample1@example.com", "Jane Doe", "This is another comment.", time.Now(), true},
    }}

    err := tmpl.ExecuteTemplate(w, "commentsIndex.html", data)
    if err != nil {
        fmt.Println(tmpl.DefinedTemplates())
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

