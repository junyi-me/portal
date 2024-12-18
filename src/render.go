package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
    tmpl = template.Must(template.ParseGlob("tmpl/*.html"))
}

func renderHome(w http.ResponseWriter) {
    tmpl.ExecuteTemplate(w, "home.html", links)
}

