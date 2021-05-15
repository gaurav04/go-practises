package main

import "fmt"

var num int

func main() {
	fmt.Println("Enter the number")
	fmt.Scan(&num)
	result := factorial(num)
	fmt.Printf("Factorial of %d is %d", num, result)
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	result := num * factorial(num-1)
	return result
}
