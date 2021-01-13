package main

type Customer struct {
		FirstName string
		LastName  string
		FullName  string
	}

func GetCustomers() (customers []Customer) {

	marcel := Customer{FirstName: "Gaurav", LastName: "Harsola"}

	customers = append(customers, marcel,Customer{FirstName: "Sanyyam", LastName: "Jain"})

	return customers

}
