package api

import (
	"net/http"
	"quizbackend/internal/api/routes"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartServer() *http.Server {
	r := mux.NewRouter()

	routes.PrepareRoutes(r)

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler: handlers.CORS(
			handlers.AllowedHeaders([]string{"content-type", "Auth-Token", ""}), //took some time to figure out middleware problem
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"}),
			handlers.AllowCredentials(),
		)(r),
	}
	return server
}
