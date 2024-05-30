package main

import (
	"log"

	"github.com/SmokierLemur51/minecraft-wms/handler"
)

func main() {
	h := handler.NewCoreHandler()

	log.Println(h.Run())
}
