package main

import (
	"encoding/json"
	"example.com/marshalling/entities"
	"fmt"
)

func main() {

	product := entities.Product{
		Id:       "p01",
		Name:     "name 1",
		Price:    4.5,
		Quantity: 8,
		Status:   true,
		Category: entities.Category{
			Id:   "c1",
			Name: "Category 1",
		},
		Comments: []entities.Comment{
			entities.Comment{Title: "title 1", Content: "content 1"},
			entities.Comment{Title: "title 2", Content: "content 2"},
			entities.Comment{Title: "title 3", Content: "content 3"},
		},
	}
	result, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}

}
