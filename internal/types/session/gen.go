package session

import (
	"database/sql"
	"math/rand"
	"quizbackend/internal/utils/db"
	"quizbackend/internal/utils/errors"
)

const letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func CheckToken(token string) (*Session, error) { //check token from auth header
	var user Session
	err := db.Db.QueryRow("SELECT UserId FROM tokens WHERE Token = $1", token).Scan(&user.Id)
	if err != nil && err == sql.ErrNoRows {
		return nil, errors.ErrorInvalidToken
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func GenToken(id int) (*Session, error) { //look back at peformance for this algorithm later
	var user Session
	user.Id = id
	err := db.Db.QueryRow("SELECT Token FROM tokens WHERE UserId = $1", id).Scan(&user.Token)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		authToken := createRandString()
		user.Token = authToken
		_, err = db.Db.Exec("INSERT INTO tokens (UserId, Token) VALUES ($1, $2)", id, authToken)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func createRandString() string {
	b := make([]byte, 128)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
