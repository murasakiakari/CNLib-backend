package controller

import (
	"CNLib-backend/utility"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


var UserMap map[string]string = make(map[string]string)

func HashPassword(password string) (string, error) {
	hashPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashPasswordByte), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func init() {
	randomPassword := utility.GetToken()
	fmt.Println("user Password: " + randomPassword)
	userPassword, _ := HashPassword(randomPassword)
	UserMap["user"] = userPassword
}
