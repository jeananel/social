package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jeananel/social.git/bd"
)

/*ViewProfile get values of user profile */
func ViewProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	profile, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(w, "An error occurred in view profile user.  "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
