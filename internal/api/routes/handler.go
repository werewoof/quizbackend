package routes

import (
	"quizbackend/internal/api/routes/user"

	"github.com/gorilla/mux"
)

func PrepareRoutes(r *mux.Router) {
	//API ROUTES REMOVE LINE BELOW AND REPLACE WITH r IF SUBDOMAIN USED
	apiRoute := r.PathPrefix("/api").Subrouter()
	//user route
	user.Routes(apiRoute)
}
