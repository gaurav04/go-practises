package main

import (
	"fmt"

	"example.com/nested-structs/entities"
)

func main() {
	var products = []entities.Product{
		entities.Product{
			Id:       "p01",
			Name:     "tivi 1",
			Price:    5,
			Quantity: 9,
			Status:   false,
			Category: entities.Category{
				Id:   "c1",
				Name: "Category 1",
			},
			Manufacturer: entities.Manufacturer{
				Id:      "m1",
				Name:    "Manufacturer 1",
				Address: "Address 1",
			},
		},
		entities.Product{
			Id:       "p02",
			Name:     "tivi 2",
			Price:    2,
			Quantity: 8,
			Status:   true,
			Category: entities.Category{
				Id:   "c1",
				Name: "Category 1",
			},
			Manufacturer: entities.Manufacturer{
				Id:      "m1",
				Name:    "Manufacturer 1",
				Address: "Address 1",
			},
		},
		entities.Product{
			Id:       "p03",
			Name:     "laptop 3",
			Price:    11,
			Quantity: 7,
			Status:   false,
			Category: entities.Category{
				Id:   "c2",
				Name: "Category 2",
			},
			Manufacturer: entities.Manufacturer{
				Id:      "m2",
				Name:    "Manufacturer 2",
				Address: "Address 2",
			},
		},
	}

	fmt.Println("Product List")
	DisplayProductList(products)
}

func DisplayProductList(products []entities.Product) {
	for _, product := range products {
		DisplayProduct(product)
		fmt.Println("-------------------------------")
	}
}

func DisplayProduct(product entities.Product) {
	fmt.Println("id: ", product.Id)
	fmt.Println("name: ", product.Name)
	fmt.Println("status: ", product.Status)
	fmt.Println("price: ", product.Price)
	fmt.Println("quantity: ", product.Quantity)
	fmt.Println("total: ", (product.Price * float64(product.Quantity)))
	fmt.Println("Category Info")
	fmt.Println("\tid: ", product.Category.Id)
	fmt.Println("\tname: ", product.Category.Name)
	fmt.Println("Manufacturer Info")
	fmt.Println("\tid: ", product.Manufacturer.Id)
	fmt.Println("\tname: ", product.Manufacturer.Name)
	fmt.Println("\taddress: ", product.Manufacturer.Address)
}
