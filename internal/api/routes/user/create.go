package user

import (
	"encoding/json"
	"io"
	"net/http"
	"quizbackend/internal/types/session"
	"quizbackend/internal/utils/db"
	"quizbackend/internal/utils/logger"
)

func create(w http.ResponseWriter, r *http.Request) {
	logger.DebugLogger.Println("user: user submited create request")

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

	var exists bool
	err = db.Db.QueryRow("SELECT EXISTS (SELECT username FROM users WHERE username = $1)", body.Username).Scan(&exists)
	if err != nil {
		logger.WarnLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists {
		logger.InfoLogger.Println("username already exists")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//hash and salt pass
	hashedpass := hashAndSalt(body.Password, body.Username)

	//insert into db
	var id int
	err = db.Db.QueryRow("INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING userid", body.Username, hashedpass, body.Email).Scan(&id)
	if err != nil {
		logger.WarnLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
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

	logger.DebugLogger.Println("user: user created")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
