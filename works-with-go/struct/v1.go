/* A struct is just a named collection of fields where you can store data. */

/* A method is a function with a receiver. Methods are very similar to functions but they are called by invoking them on an instance of a particular type.
   Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things". */

/* Interfaces are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types  */
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area returns the area of the triangle.
func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}

func main() {
	var s Shape = Rectangle{10, 12}
	var s1 Shape = Circle{10}
	var s2 Shape = Triangle{10, 12}
	fmt.Println(s.Area())
	fmt.Println(s1.Area())
	fmt.Println(s2.Area())
}
