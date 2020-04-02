package main

import (
	"fmt"

	"example.com/polymorphisam/entities"
)

func main() {

	circle1 := entities.Circle{Name: "Circle 1", R: 4}
	circle2 := entities.Circle{Name: "Circle 2", R: 7}
	rectangle1 := entities.Rectangle{Name: "Rectangle 1", A: 4, B: 10}
	rectangle2 := entities.Rectangle{Name: "Rectangle 1", A: 2, B: 7}
	square1 := entities.Square{Name: "Square 1", A: 6}
	square2 := entities.Square{Name: "Square 1", A: 10}

	geometries := []entities.Geometry{circle1, rectangle1, square1, rectangle2, square2, circle2}

	for _, geometry := range geometries {
		fmt.Println(geometry.GetName())
		fmt.Printf("Area: %0.2f\n", geometry.Area())
		fmt.Printf("Perimeter: %0.2f\n", geometry.Perimeter())
		fmt.Println("----------------------")
	}

}
