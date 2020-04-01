package main

import (
	"encoding/json"
	"example.com/marshalling/entities"
	"fmt"
)

func main() {

	var str string = `{
		"id":"p01",
		"name":"name 1",
		"price":4.5,
		"quantity":8,
		"status":true,
		"category":{
			"id":"c1",
			"name":"Category 1"
		},
		"comments":[
			{"title":"title 1","content":"content 1"},
			{"title":"title 2","content":"content 2"},
			{"title":"title 3","content":"content 3"}
		]
	}`

	var product entities.Product
	json.Unmarshal([]byte(str), &product)

	fmt.Println("Product Info")
	fmt.Printf("Id: %s\nName: %s\nPrice: %0.2f\nQuantity: %d\nStatus: %t", product.Id, product.Name, product.Price, product.Quantity, product.Status)
	fmt.Println("Category Info")
	fmt.Println("\tCategory Id: ", product.Category.Id)
	fmt.Println("\tCategory Name: ", product.Category.Name)
	fmt.Println("Comment List")
	for _, comment := range product.Comments {
		fmt.Println("\t", comment.Title)
		fmt.Println("\t", comment.Content)
		fmt.Println("\t-----------------------")
	}

}
