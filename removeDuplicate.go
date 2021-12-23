package main

import "fmt"

func removeDuplicate(arr []int) []int {
	result := []int{}
	occurred := make(map[int]bool)
	for _, e := range arr {
		if occurred[arr[e]] != true {
			occurred[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 4, 5, 2}
	fmt.Println(removeDuplicate(arr))
}
