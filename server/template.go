package server 

import (
	"net/http"
	"html/template"
)

type TemplateData struct {
	Page     string
	Title    string
	Content  string
	CSS      string
}

var CSS_URL string = "/static/css/testing.css"

func (t TemplateData) RenderHTMLTemplate(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/" + t.Page)
	if err != nil {
		return
	}
	err = tmpl.Execute(w, t)
	if err != nil {
		return
	}
	// this prevents the superflous hanlder err
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
