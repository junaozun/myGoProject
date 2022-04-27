package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	con, err := redis.Dial("tcp", "127.0.0.1:9898")
	if err != nil {
		fmt.Println("redis conn fail：", err)
		return
	}
	defer con.Close()
	_, err = con.Do("flushdb")
	if err != nil {
		fmt.Println("flusdhdb err：", err)
		return
	}
	fmt.Println("success")
}
