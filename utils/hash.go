package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hashing(pass string) ([]byte, error) {
	salt := 8
	fmt.Println(pass)
	return bcrypt.GenerateFromPassword([]byte(pass), salt)
}

func CompareHash(hash, pass string) error {
	fmt.Println(pass)
	fmt.Println([]byte(hash), []byte(pass))
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	fmt.Println(err)
	return err
}
