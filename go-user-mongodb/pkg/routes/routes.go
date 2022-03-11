package routes

import (
	"github.com/go-user-mongodb/pkg/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")
	return router
}
