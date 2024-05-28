package main

import (
  "os"
  "log"
  "net/http"
  "github.com/SmokierLemur51/leedev/server"
)

func init() {
  os.SetEnv("database", "testing")
}

func main() {
  s := server.Server
  s.ConfigServer()
  
  log.Println("Listening on port ", s.Port)
  log.Fatal(http.ListenAndServe(s.Port, s.Router))
}
