package main

import (
	"fmt"

	"example.com/interface/entities"
)

func main() {

	var animal entities.Animal
	dog := entities.Dog{}
	cat := entities.Cat{}

	animal = dog
	fmt.Println(animal.Sound())

	animal = cat
	fmt.Println(cat.Sound())

}
