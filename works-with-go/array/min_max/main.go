package main

import "fmt"

func main() {
	var a [100]int
	var num int

	fmt.Println("Enter the number you want in Array")
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		fmt.Printf("Enter the %d Element :", i)
		fmt.Scan(&a[i])
	}
	get_max(a, num)
	fmt.Println("")
	get_min(a, num)

}

func get_max(arr [100]int, num int) {
	large := arr[0]
	for i := 1; i < num; i++ {
		if large < arr[i] {
			large = arr[i]
		}
	}
	fmt.Printf("The largest element is %d", large)
}

func get_min(arr [100]int, num int) {
	small := arr[0]
	for i := 1; i < num; i++ {
		if small > arr[i] {
			small = arr[i]
		}
	}
	fmt.Printf("The Smallest element is %d", small)
}
