package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World")

	// array declartion
	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	// Slices
	weekdays := days[0:5]
	fmt.Println(weekdays)

	//Maps
	youtubeSubscribers := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fun Fun Function": 171220,
	}
	fmt.Println(youtubeSubscribers["MKBHD"])

	//Struct
	type Person struct {
		name string
		age  int
	}

	person := Person{name: "Elliot", age: 24}
	fmt.Println(person.name)

	type Route struct {
		Name    string
		Method  string
		Pattern string
	}

	// Routes defines the list of routes of our API
	var Routes []Route

	Routes = []Route{
		Route{"Authentication", "POST", "/get-token"},
		Route{"Authentication", "POST", "/get-token"},
	}

	fmt.Println(Routes)

}
