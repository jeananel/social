package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jeananel/social.git/bd"
	"github.com/jeananel/social.git/jwt"
	"github.com/jeananel/social.git/models"
)

//Login for user access to aplication trought json web token
func Login(write http.ResponseWriter, request *http.Request) {
	write.Header().Add("content-type", "application/json")

	var object models.User

	err := json.NewDecoder(request.Body).Decode(&object)

	if err != nil {
		http.Error(write, "Invalid user or password. "+err.Error(), 400)
		return
	}

	//Validations
	if len(object.Email) == 0 {
		http.Error(write, "Email is required.", 400)
		return
	}

	document, loginOK := bd.TryLoginUser(object.Email, object.Password)

	if !loginOK {
		http.Error(write, "Invalid user or password.", 400)
		return
	}

	tokenGenerated, err := jwt.GenerateJWT(document) //Token generated
	if err != nil {
		http.Error(write, "An error occurred while generating the token."+err.Error(), 400)
		return
	}

	responseToken := models.Token{
		Token: tokenGenerated,
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)
	json.NewEncoder(write).Encode(responseToken)

	//Set token to cookies
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(write, &http.Cookie{
		Name:    "token",
		Value:   tokenGenerated,
		Expires: expirationTime,
	})

}
