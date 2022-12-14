package user

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd string, username string) string {
	byteString := []byte(username + pwd)
	hash, err := bcrypt.GenerateFromPassword(byteString, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(pwd string, userHashedPwd string) bool {
	byteHash := []byte(pwd)
	byteUserHash := []byte(userHashedPwd)
	err := bcrypt.CompareHashAndPassword(byteUserHash, byteHash)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
