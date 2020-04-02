package main

import (
	"example.com/pointer-rec/entities"
	"fmt"
)

func main() {
	product1 := entities.Product{
		Id:       "p01",
		Name:     "tivi 1",
		Price:    5,
		Quantity: 9,
		Status:   false,
	}

	product1.ChangeValue1("abc")
	fmt.Println("Product 1 Info")
	fmt.Println(product1.ToString())

	product2 := entities.NewProduct("p02", "name 2", 2, 7, true)
	product2.ChangeValue2("def")
	fmt.Println("Product 2 Info")
	fmt.Println(product2.ToString())
}
