package user

import "github.com/gorilla/mux"

func Routes(r *mux.Router) {
	//user route
	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/auth", auth).Methods("POST")
	user.HandleFunc("/create", create).Methods("POST")
}
