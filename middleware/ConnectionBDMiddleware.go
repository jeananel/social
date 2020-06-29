package middleware

import (
	"net/http"

	"github.com/jeananel/social.git/bd"
)

//CheckConnectionToDatabase middleware for verify connection to database
func CheckConnectionToDatabase(next http.HandlerFunc) http.HandlerFunc {

	return func(write http.ResponseWriter, request *http.Request) {
		if !bd.CheckConnection() {
			http.Error(write, "Dont establishing connection to database", 500)
			return
		}
		next.ServeHTTP(write, request)
	}
}
