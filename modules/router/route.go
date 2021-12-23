package router

import (
	"golang-rest-api/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook).Methods("GET")
	return router
}
