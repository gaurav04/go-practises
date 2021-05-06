package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Phonedetails struct {
	Home   string `json:"home,omitempty"`
	Office string `json:"office,omitempty"`
}

type Details struct {
	Email    string       `json:"email,omitempty"`
	Jobtitle string       `json:"jobtitle,omitempty"`
	Name     string       `json:"name,omitempty"`
	Phone    Phonedetails `json:"phone,omitempty"`
}

var details Details

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	file, err := ioutil.ReadFile("a.json")
	checkError(err)

	err = json.Unmarshal(file, &details)
	checkError(err)

	fmt.Println(details.Email)
	fmt.Println(details.Jobtitle)
	fmt.Println(details.Name)
	fmt.Println(details.Phone.Home)
	fmt.Println(details.Phone.Office)
}
