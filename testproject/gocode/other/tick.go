package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		t := time.NewTicker(time.Second * 5)
		for {
			<-t.C
			fmt.Println("iiii")
		}
	}()
	for {

	}
}
