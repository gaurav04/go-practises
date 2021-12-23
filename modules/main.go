package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-rest-api/router"
)

func main() {

	fmt.Println("Welcome to test API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
