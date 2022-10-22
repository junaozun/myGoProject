package main

import (
	"fmt"
	"io"
	"net"

	"github.com/xtaci/kcp-go"
)

func main() {
	fmt.Println("kcp listens on 10000")
	lis, err := kcp.ListenWithOptions(":10000", nil, 10, 3)
	if err != nil {
		panic(err)
	}
	for {
		conn, e := lis.AcceptKCP()
		if e != nil {
			panic(e)
		}
		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	var buffer = make([]byte, 1024, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}
		fmt.Println("receive from client:", string(buffer[:n]))
	}
}
