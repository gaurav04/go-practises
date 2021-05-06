package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const (
	Port = ":8080"
)

type Details struct {
	Country string
	City    string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var details Details

	details.Country = "India"
	details.City = "Pune"

	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, details)
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(Port, nil)
}
