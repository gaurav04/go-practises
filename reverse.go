package main

import "fmt"

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {

	arr := []int{1, 2, 3, 6, 4}
	reverse(arr)
	fmt.Println(arr)
}
