package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jeananel/social.git/bd"
	"github.com/jeananel/social.git/models"
)

/*ViewProfile get values of user profile */
func ViewProfile(write http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(write, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	profile, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(write, "An error occurred in view profile user.  "+err.Error(), 400)
		return
	}

	write.Header().Set("context-type", "application/json")
	write.WriteHeader(http.StatusCreated)
	json.NewEncoder(write).Encode(profile)
}

/*UpdateProfile update user profile */
func UpdateProfile(write http.ResponseWriter, request *http.Request) {

	var t models.User

	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(write, "Invalid data "+err.Error(), 400)
		return
	}

	var status bool

	//IDUser is global variable in Token
	status, err = bd.UpdateUser(t, IDUser)
	if err != nil {
		http.Error(write, "An error occurred in update profile. Try again please. "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(write, "An error occurred in update profile. ", 400)
		return
	}

	write.WriteHeader(http.StatusCreated)

}
