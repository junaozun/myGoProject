package main

import (
	"fmt"
)

type IDrawStrategy interface {
	IDraw(fn func([]int, int) string, sa func(string)) error
}

type ihm struct {
}

func (m ihm) IDraw(fn func([]int, int) string, sa func(string)) error {
	b := fn([]int{3, 4}, 1)
	fmt.Println(b)
	sa("nihao")
	return nil
}

func (m ihm) dag() {
	fmt.Println("999999")
	err := m.IDraw(func(ints []int, dd int) string {
		if dd == 1 {
			return ""
		}
		fmt.Println("ooooooo")
		var kk string
		for _, v := range ints {
			kk += fmt.Sprint(v)
		}
		kk += fmt.Sprint(dd)
		return kk
	}, func(s string) {
		fmt.Println(s)
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	oe := &ihm{}
	oe.dag()
}
