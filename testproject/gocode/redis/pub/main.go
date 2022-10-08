package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	var count int
	for {
		count++
		if count >= 10 {
			client.Publish("channel1", "close").Result()
			break
		}
		time.Sleep(time.Second * 2)
		ret, err := client.Publish("channel1", "message").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(ret)
	}
}
