package main

import (
	"log"

	"github.com/jeananel/social/bd"
	"github.com/jeananel/social/handlers"
)

func main() {

	if bd.CheckConnection() {
		log.Fatal("Dont connect. Error.")
		return
	}

	handlers.Managements()
}
