package main

import "fmt"

func main() {

	var number, remainder, temp int

	var reverse int = 0

	fmt.Print("Enter the number")
	fmt.Scan(&number)

	temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number = number / 10

		if number == 0 {
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d is a Palindrome", temp)
	} else {
		fmt.Printf("%d is not a Palindrome", temp)
	}

}
