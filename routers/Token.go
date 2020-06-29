package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeananel/social.git/bd"
	"github.com/jeananel/social.git/models"
)

/*Email email value for endpoint */
var Email string

/*IDUser ID in model for endpoints */
var IDUser string

/*VerifyToken get values  token */
func VerifyToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("Devjp")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Invalid format token")
	}

	token = strings.TrimSpace(splitToken[1])

	resultToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := bd.CheckExistUser(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !resultToken.Valid {
		return claims, false, "", errors.New("Invalid token")
	}
	return claims, false, "", err
}
