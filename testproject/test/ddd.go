package main

import (
	"fmt"
)

type suxf interface {
	nihao()
	say()
}

type aaa interface {
	ben()
}

type zong interface {
	suxf
	aaa
}

type one struct {
}

func (o one) nihao() {
	fmt.Println("nihao")
}

func (o one) say() {
	fmt.Println("say")
}

func (o one) ben() {
	fmt.Println("ben")
}

func main() {
	t := one{}
	var s suxf
	s = t
	if cmd, ok := s.(zong); ok {
		cmd.nihao()
		cmd.say()
		cmd.ben()
	}
	fmt.Println("cc")
}
