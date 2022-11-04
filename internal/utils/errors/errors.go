package errors

import "errors"

var (
	//user route
	ErrorInvalidUserDetails = errors.New("user: invalid user details")
	ErrorUserAlreadyExists  = errors.New("user: user already exists")

	//session
	ErrorInvalidToken = errors.New("session: invalid token")
)
