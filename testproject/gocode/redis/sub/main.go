package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	var wg sync.WaitGroup
	// 这里我们用携程来模拟多个服务
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			pb := client.Subscribe("channel1")
			fmt.Println("阻塞，等待读取 Channel 信息")
			for {
				select {
				case mg := <-pb.Channel():
					// 等待从 channel 中发布 close 关闭服务
					if mg.Payload == "close" {
						// 当
						wg.Done()
					} else {
						log.Println("接channel信息", mg.Payload)
					}
				default:
				}
			}
		}(i)
	}
	wg.Wait()
	log.Println("结束 channel")
}
