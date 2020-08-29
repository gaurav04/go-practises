package main

import (
	"net/http"

	"github.com/gaurav04/go-practises/inventroy_mgmt/database"
	"github.com/gaurav04/go-practises/inventroy_mgmt/product"
)

func main() {
	database.SetupDatabase()
	product.SetupRoutes()
	http.ListenAndServe(":5002", nil)
}
