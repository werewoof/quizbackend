package user

import (
	"encoding/json"
	"io"
	"net/http"
	"quizbackend/internal/types/session"
	"quizbackend/internal/utils/db"
	"quizbackend/internal/utils/errors"
	"quizbackend/internal/utils/logger"
)

func auth(w http.ResponseWriter, r *http.Request) {
	//do something
	logger.DebugLogger.Println("user: user submited auth request")
	//gather data from request
	var body account
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.InfoLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		logger.InfoLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//validate
	err = validateInput(body)
	if err != nil {
		logger.InfoLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//hash and salt pass
	pass := body.Username + body.Password
	//select from db
	var id int
	var userHashedPass string
	err = db.Db.QueryRow("SELECT UserId, Password FROM users WHERE username = $1", body.Username).Scan(&id, &userHashedPass)
	if err != nil {
		logger.InfoLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//compare pass
	if !comparePasswords(pass, userHashedPass) {
		logger.InfoLogger.Println(errors.ErrorInvalidUserDetails)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := session.GenToken(id)

	if err != nil {
		logger.WarnLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(user)
	if err != nil {
		logger.WarnLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.DebugLogger.Println("user: user logged in")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
