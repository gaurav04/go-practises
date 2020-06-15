package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}

	type Social struct {
		Facebook string `json:"facebook"`
		Twitter  string `json:"twitter"`
	}

	type User struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Age    int    `json:"Age"`
		Social Social `json:"social"`
	}

	type Users struct {
		Users []User `json:"users"`
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

}
