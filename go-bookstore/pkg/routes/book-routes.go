package routes

import (
	"github.com/go-bookstore/pkg/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/book", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controller.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
	return router
}
