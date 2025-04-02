package server

import (
	"forum/internal/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHanlder).Methods("POST")
	r.HandleFunc("/regiser", handlers.RegiserHandler).Methods("POST")
	return r
}
