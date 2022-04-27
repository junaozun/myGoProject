package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

var redisUrl = flag.String("redisUrl", "", "redisUrl")
var redisIndex = flag.String("index", "", "redis index")

func main() {
	flag.Parse()

	ri, err := strconv.Atoi(*redisIndex)
	if err != nil {
		fmt.Println("redis index convert err：", err)
		return
	}
	if ri > 100 || ri < 0 {
		fmt.Println("redis index out of range,need 0-100：", err)
		return
	}

	//1.连接到redis数据库
	con, err := redis.Dial("tcp", *redisUrl)
	if err != nil {
		fmt.Println("redis conn fail：", err)
		return
	}
	defer con.Close()
	fmt.Println("redis connection success")

	_, err = con.Do("auth", "sanguo")
	if err != nil {
		fmt.Println("auth err：", err)
		return
	}

	_, err = con.Do("select", *redisIndex)
	if err != nil {
		fmt.Println("select db err：", err)
		return
	}
	fmt.Println("select db success")

	_, err = con.Do("flushdb")
	if err != nil {
		fmt.Println("flusdhdb err：", err)
		return
	}
	fmt.Println("flushdb seccess")
}
