package utils

import (
	"golang.org/x/crypto/bcrypt"
)

//EcryptPasswordUtil for encript password
func EcryptPasswordUtil(password string) (string, error) {
	cost := 8 // 2^8 for method encriptation
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
