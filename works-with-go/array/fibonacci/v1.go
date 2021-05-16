package main

import "fmt"

var num int

func main() {
	fmt.Println("Enter the number")
	fmt.Scan(&num)
	fmt.Println("Fibonacci series ....")
	for i := 0; i < num; i++ {
		fmt.Print(fibonacci(i), ", ")
	}
}

func fibonacci(num int) int {
	if num == 0 || num == 1 {
		return num
	} else {
		return (fibonacci(num-1) + fibonacci(num-2))
	}
}
