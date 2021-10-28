package db

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	// salt := 6 // common users so so
	salt := 8 // super users

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), salt)

	return string(bytes), err
}
