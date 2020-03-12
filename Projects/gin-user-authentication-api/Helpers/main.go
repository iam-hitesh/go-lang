package Helpers


import (
	"golang.org/x/crypto/bcrypt"
)


func GetHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 9)

	return string(bytes), err
}


func VerifyPassword(hashedPassword, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))

	return err == nil
}