package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeananel/social.git/models"
)

//GenerateJWT for generate json web token in application
func GenerateJWT(user models.User) (string, error) {

	mainKey := []byte("Devjp")

	payload := jwt.MapClaims{
		"Email":     user.Email,
		"Name":      user.Name,
		"LasName":   user.LastName,
		"DateBirth": user.DateBirth,
		"Biography": user.Biography,
		"Location":  user.Location,
		"WebSite":   user.WebSite,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenResult, err := token.SignedString(mainKey)
	if err != nil {
		return tokenResult, err
	}

	return tokenResult, nil
}
