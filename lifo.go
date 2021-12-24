package main

import "fmt"

func main() {

	var stack []int

	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)

	for i := len(stack) - 1; i >= 0; i-- {

		fmt.Println(stack[i])
	}
}
