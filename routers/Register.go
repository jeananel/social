package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jeananel/social.git/bd"
	"github.com/jeananel/social.git/models"
)

//Register for register an user
func Register(write http.ResponseWriter, request *http.Request) {
	var object models.User
	err := json.NewDecoder(request.Body).Decode(&object)
	if err != nil {
		http.Error(write, "An error ocurred in user register. "+err.Error(), 400)
		return
	}

	//Validations
	if len(object.Email) == 0 {
		http.Error(write, "Email is required.", 400)
		return
	}
	if len(object.Password) < 6 {
		http.Error(write, "Password invalid, must be at least 6 characters.", 400)
		return
	}

	_, userFounded, _, err := bd.CheckExistUser(object.Email)

	if userFounded {
		http.Error(write, "The email has already been registered."+err.Error(), 400)
		return
	}

	_, status, err := bd.InsertRegister(object)

	if err != nil {
		http.Error(write, "An error occurred in insert register user."+err.Error(), 400)
		return
	}

	if !status {
		http.Error(write, "Not insert user register."+err.Error(), 400)
		return
	}

	write.WriteHeader(http.StatusCreated)
}
