package main

import (
	"fmt"
	"net"
)

func main() {
	cliV1()
}

func cliV1() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	if _, err := conn.Write([]byte("hello")); err != nil {
		fmt.Println(err)
	}

	buffer := make([]byte, 1024)
	recvNum, err := conn.Read(buffer)

	msg := string(buffer[:recvNum])
	fmt.Println("recv from server: ", msg)
}

//解决问题4
func cliV4() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	if _, err := conn.Write([]byte("hello")); err != nil {
		fmt.Println(err)
	}

	buffer := make([]byte, 1024)
	recvNum, err := conn.Read(buffer)

	msg := string(buffer[:recvNum])
	fmt.Println("recv from server: ", msg)

	req := []byte("hello")
	sendNum := 0
	num := 0
	// 循环发包
	for sendNum < len(req) {
		num, err = conn.Write(req[sendNum:])
		if err != nil {
			fmt.Println(err)
			break
		}
		sendNum += num
	}

	recvNum, err = conn.Read(buffer)

	msg = string(buffer[:recvNum])
	fmt.Println("recv from server: ", msg)
}

// 解决问题5，使用客户端连接池的方式
