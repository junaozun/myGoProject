package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10, 32}
	if len(arr2) > len(arr1) {
		arr2 = arr2[:len(arr1)]
	}
	fmt.Println(arr1, arr2)
}
