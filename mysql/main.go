package main

import (
	"net/http"

	"github.com/gaurav04/go-practises/mysql/database"
	"github.com/gaurav04/go-practises/mysql/product"
)

func main() {
	database.SetupDatabase()
	product.SetupRoutes()
	http.ListenAndServe(":5002", nil)
}
