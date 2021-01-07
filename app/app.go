package main

import (
	"fmt"
)

func main() {

	customers := GetCustomers()

	for _, customer := range customers {
		fmt.Println(customer)
	}

}

func getData() (customers []string) {

	customers = []string{"Gaurav Harsola", "Golu Jain"}

	customers = append(customers, "Sanyyam Jain")

	for _, customer := range customers {

		fmt.Println(customer)
	}

	return customers
}
