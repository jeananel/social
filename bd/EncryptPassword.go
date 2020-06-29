package bd

import "golang.org/x/crypto/bcrypt"

//EcryptPass for encript password
func EcryptPass(password string) (string, error) {
	cost := 8 // 2^8 for method encriptation
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
