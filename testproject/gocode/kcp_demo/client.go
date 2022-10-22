package main

import (
	"strconv"
	"time"

	"github.com/xtaci/kcp-go"
)

func main() {
	kcpconn, err := kcp.DialWithOptions("localhost:10000", nil, 10, 3)
	if err != nil {
		panic(err)
	}

	var count int
	for {
		count++
		kcpconn.Write([]byte("hello suxf" + strconv.Itoa(count)))
		time.Sleep(5 * time.Second)
	}
}
