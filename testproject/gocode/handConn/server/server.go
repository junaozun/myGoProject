package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	ServerV1()
}

func ServerV1() {
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	conn, err := lis.Accept()
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 1024)
	recvNum, err := conn.Read(buffer)

	msg := string(buffer[:recvNum])
	fmt.Println("recv from client: ", msg)

	conn.Write([]byte("world"))

}




















/*----------------------------------------------------------------------------------------------------------------------*/
// 解决问题1,使 server 的读写异步化，这里其实使用 go 语言的协程机制很方便实现
func ServerV2() {
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.Accept()
		defer conn.Close()
		if err != nil {
			panic(err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	recvNum, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}

	msg := string(buffer[:recvNum])
	fmt.Println("recv from client: ", msg)

	conn.Write([]byte("world"))
}

/*----------------------------------------------------------------------------------------------------------------------*/
// 要解决问题2，我们可以采用长连接的方式，每次 client 和 server 建立连接后，这个连接默认存活，只有在对端关闭连接时，或者连接一直空闲、超过指定时间没有发送消息时，这个连接才关闭。
//在连接的存活期内，对对端发来的请求进行循环读写，这样就能避免短连接不断创建和关闭造成的性能损耗。这里只需要更改上面的 handleConn 这个方法即可，如下：
func ServerV3() {
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.Accept()
		defer conn.Close()
		if err != nil {
			panic(err)
		}
		go handleConn2(conn)
	}

}
func handleConn2(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		recvNum, err := conn.Read(buffer)
		if err == io.EOF {
			// client 连接关闭
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}

		msg := string(buffer[:recvNum])
		fmt.Println("recv from client: ", msg)

		conn.Write([]byte("world"))
	}
}

// 解决问题3，初始化一块内存就好了

// 解决问题4,改动 server 的 handleConn 方法，使用 io.ReadFull 进行包读取，如下：
func handleConn3(conn net.Conn) {

	for {
		buffer := make([]byte, 5)
		// 使用 io.ReadFull 进行包读取
		recvNum, err := io.ReadFull(conn, buffer)
		if err == io.EOF {
			// client 连接关闭
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}

		msg := string(buffer[:recvNum])
		fmt.Println("recv from client: ", msg)

		conn.Write([]byte("world"))
	}

}
