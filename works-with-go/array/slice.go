/* Arrays have a fixed capacity which you define when you declare the variable. We can initialize an array in two ways:
[N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
[...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5} */

package main

import "fmt"

// Sum calculates the total from an array of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	num := []int{1, 2, 3, 4, 5}
	x := Sum(num)
	fmt.Println(x)
}
