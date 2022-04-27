package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic nihao")
		}
	}()
	a := 10
	var c int
	for i := 0; i < 1000; i++ {
		c += a
		if c == 100 {
			fmt.Println(c)
			panic("anglany")
		}
	}
	fmt.Println(c)
}
