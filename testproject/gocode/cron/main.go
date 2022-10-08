package main

import (
	"fmt"
	"time"
)

// func main() {
// 	c := cron.New()
// 	c.Start()
// 	c.AddFunc("@every 1m", func() {
// 		fmt.Println("tick every 1 second")
// 	})
// 	select {}
// }
//
func main() {

	// 协程A
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("goroutine1_print")
		}
	}()

	// 协程B
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("33333333333333333")
			}
		}()
		time.Sleep(2 * time.Second)
		fmt.Println("00000000000000000000000000")
		panic("goroutine2_panic")
	}()

	time.Sleep(100 * time.Second)
}
