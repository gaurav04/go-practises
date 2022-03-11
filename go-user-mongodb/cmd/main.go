package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-user-mongodb/pkg/routes"
)

func main() {
	fmt.Println("Welcome User")
	r := routes.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")

}
