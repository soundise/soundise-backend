package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type UsersHandler struct{}

func (handler UsersHandler) New(router *mux.Router) *UsersHandler {
	// Register routes to the router
	router.HandleFunc("/api/users", handler.Add).Methods("POST")
	router.HandleFunc("/api/users", handler.Delete).Methods("DELETE")

	return &handler
}

func (handler UsersHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (handler UsersHandler) Add(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

}
