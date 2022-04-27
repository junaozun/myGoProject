package main

import (
	"fmt"
	"net"
)

func main() {
	cliV1()
}

func cliV1() {
	// go使用dial对soket链接过程进行了封装，实际上客户端调用了connect()函数与服务器建立链接
	//对于客户端的 connect() 函数，该函数的功能为客户端主动连接服务器，建立连接是通过三次握手，而这个连接的过程是由内核完成，不是这个函数完成的，这个函数的作用仅仅是通知 Linux 内核，
	//让 Linux 内核自动完成 TCP 三次握手连接，最后把连接的结果返回给这个函数的返回值（成功连接为0， 失败为-1）。
	//常的情况，客户端的 connect() 函数默认会一直阻塞，直到三次握手成功或超时失败才返回（正常的情况，这个过程很快完成）。
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
