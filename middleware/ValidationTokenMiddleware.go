package middleware

import (
	"net/http"

	"github.com/jeananel/social.git/routers"
)

/*ValidationToken validate token in requests */
func ValidationToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.VerifyToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error in token."+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
