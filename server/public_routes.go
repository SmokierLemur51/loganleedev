package server

import (
  "github.com/go-chi/chi/v5"
)

var PUBLIC_CSS string = "/static/css/testing.css"

func (s Server) IndexHandler() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    t := TemplateData{Page: "public/index.html", Title: "Test"}
    t.RenderTemplate(w)
  }
}
