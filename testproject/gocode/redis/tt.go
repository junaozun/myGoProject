package main

import (
	"fmt"
)

func main() {
	fmt.Println(UpIntegerTen(21))
}

func UpIntegerTen(num int) int {
	if num <10 {
		return 10
	}
	res := num -num %10
	if num %10 != 0 {
		res += 10
	}
	return res
}