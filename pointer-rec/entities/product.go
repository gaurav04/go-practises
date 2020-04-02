package entities

import "fmt"

type Product struct {
	Id       string
	Name     string
	Price    float64
	Quantity int
	Status   bool
}

func NewProduct(id string, name string, price float64, quantity int, status bool) Product {
	product := Product{id, name, price, quantity, status}
	return product
}

func (product Product) ToString() string {
	return fmt.Sprintf("id: %s\nname: %s\nprice: %0.2f\nquantity: %d\nstatus: %t", product.Id, product.Name, product.Price, product.Quantity, product.Status)
}

func (product Product) ChangeValue1(name string) {
	product.Name = name
}

func (product *Product) ChangeValue2(name string) {
	product.Name = name
}
