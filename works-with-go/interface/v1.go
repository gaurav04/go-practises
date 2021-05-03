/* Any type that implements the Alert() method is said to satisfy the Vehicle interface.

When defining a struct you do not explicitly specify the interface it implements.
Go determines automatically if a given type satisfies an interface if it defines all the methods. */

package main

import "fmt"

type Vehicle interface {
	Alert() string
}

type Car struct{}

func (c Car) Alert() string {
	return "Honk! Honk!"
}

type Bicycle struct{}

func (b Bicycle) Alert() string {
	return "Ring! Ring!"
}

func main() {
	c := Car{}
	b := Bicycle{}

	vehicles := []Vehicle{c, b} //create an array of Vehicles
	for _, v := range vehicles {
		fmt.Println(v.Alert())
	}
}
