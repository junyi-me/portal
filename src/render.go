package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
    tmpl = template.Must(template.ParseGlob("tmpl/*.html"))
}

type HomeData struct {
    Hosted   []Link
    External []Link
}
var homeData = HomeData{hosted, external}

func renderHome(w http.ResponseWriter) {
    tmpl.ExecuteTemplate(w, "home.html", homeData)
}

