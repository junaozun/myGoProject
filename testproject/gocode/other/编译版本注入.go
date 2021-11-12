package main

import "fmt"

var version = 1.0

func main() {
	fmt.Println("版本注入")
}

/*
go build -ldflags "-X main.version=1.2" -o 编译版本注入.go
*/
