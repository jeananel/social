package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jeananel/social.git/middleware"
	"github.com/jeananel/social.git/routers"
	"github.com/rs/cors"
)

//Managements for set server and route listener ports
func Managements() {
	router := mux.NewRouter()

	//Routes
	router.HandleFunc("/Register", middleware.CheckConnectionToDatabase(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	//Access to any users to API
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
