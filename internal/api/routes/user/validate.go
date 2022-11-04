package user

import (
	"quizbackend/internal/utils/errors"
	"regexp"
)

type account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func validateInput(body account) error {
	if body.Username == "" || body.Password == "" || body.Email == "" {
		return errors.ErrorInvalidUserDetails
	}
	usernameValid, err := regexp.MatchString("^[a-zA-Z0-9_]{3,20}$", body.Username)
	if err != nil {
		return err
	}
	if !usernameValid {
		return errors.ErrorInvalidUserDetails
	}
	return nil
}
