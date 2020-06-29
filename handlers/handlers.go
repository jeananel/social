package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Managements for set server and route listener ports
func Managements() {
	router := mux.NewRouter()
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	//Access to any users to API
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
