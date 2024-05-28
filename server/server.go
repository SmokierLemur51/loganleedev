package server

import (
  "fmt"
  "net/http"
  "os"
  "github.com/go-chi/chi/v5"
  "github.com/jmoiron/sqlx"
  _ "github.com/mattn/go-sqlite3"
)

type Server struct {
  Port   string
  Router *chi.Mux
  DB     *sqlx.DB
}

func (s *Server) ConfigServer() {
  s.Port = ":5000"
  // setup chi
  s.Router = chi.NewRouter()

  // database connection 
  var e error
  s.DB, e = sqlx.Connect("sqlite", fmt.Sprintf("instance/%s.db", os.GetEnv("database")))
  if e != nil {
    panic(e)
  }

  // static
  s.Router.Handle("/static/*", 
    http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}



func (s Server) GroupPublicRoutes() {
  s.R.Group(func(r chi.Router){
    r.Method(http.MethodGet, "/", s.IndexHandler)
  })
}

func (s Server) GroupPortalRoutes() {}
