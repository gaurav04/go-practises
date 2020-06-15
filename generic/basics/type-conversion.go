package main

import (
	"fmt"
	"strconv"
)

func main() {
	strToIntConversion()
	intToStringConversion()
}

func strToIntConversion() {
	fmt.Println("String to Integer Value Conversion")

	ourInteger, err := strconv.Atoi("12345")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ourInteger + 1)
}

func intToStringConversion() {
	fmt.Println("integer to string conversion")

	var ourString string

	ourString = strconv.Itoa(12345)

	fmt.Println(ourString)

}
