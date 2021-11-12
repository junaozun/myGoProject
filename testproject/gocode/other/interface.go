package main

import "fmt"

type sxf interface {
	say(a int)
	hello() int
}

type aa struct {
	a int
}

func (ms aa) say(a int) {
	fmt.Println("aa", a)
}

func (ms aa) hello() int {
	fmt.Println("bbb")
	return 0
}

type bb struct {
	b int
}

func (ff bb) say(a int) {
	fmt.Println("bb", a)
}

func (ff bb) hello() int {
	fmt.Println("bbb")
	return 0
}

func (ff bb) name() {

}

func main() {
	aaa := &aa{
		a: 1032,
	}
	bbb := &bb{
		b: 2021,
	}

	runLoginc := []sxf{aaa}
	runLoginc = append(runLoginc, bbb)
	fmt.Printf("%v", runLoginc)

	for _, v := range runLoginc {
		v.say(10)
		v.hello()
	}
	fmt.Println(runLoginc)
}
