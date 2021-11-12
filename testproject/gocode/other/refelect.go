package main

import (
	"log"
	"reflect"
	"strings"
)

type Hello struct {
	Sxf string
	Id  int
}

func (h *Hello) Test1(a int) {
	h.Id = a
}
func (h *Hello) Test2(b string) {
	h.Sxf = b
}

func main() {
	var hello Hello
	typ := reflect.TypeOf(&hello)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		// j 从 1 开始，第 0 个入参是 hello 自己。
		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}
}
