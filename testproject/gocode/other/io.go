package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
)

func main() {
	var conn net.Conn
	// 从链接conn中的数据一次性全部读出来
	b, _ := ioutil.ReadAll(conn)
	fmt.Println(string(b))

	buf := make([]byte, 10)

	// 从链接中读取<= len(buf) 长度的数据,读取出来的数据存入buf中
	n, _ := conn.Read(buf)
	fmt.Println(n, string(buf))

	// 从链接中必须(至少)读取len(buf) 长度的数据，读取出来的数据存入buf中，如果没读出len（buf）长度的数据，报ErrUnexpectedEOF说明数据有问题，被修改了
	nn, _ := io.ReadFull(conn, buf)
	fmt.Println(nn, string(buf))
}
