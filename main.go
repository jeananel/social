package main

import (
	"log"

	"github.com/jeananel/social.git/bd"
	"github.com/jeananel/social.git/handlers"
)

func main() {

	if !bd.CheckConnection() {
		log.Fatal("Dont connect. Error.")
		return
	}

	handlers.Managements()
}
