package main

//
// 引用的C头文件需要在注释中声明，紧接着注释需要有import "C"，且这一行和注释之间不能有空格
//

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"


import (
	"fmt"
	"unsafe"
)

func main() {
	//使用C.CString创建的字符串需要手动释放。
	cs := C.CString("Hello World\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
	fmt.Println("call C.sleep for 3s")
	C.sleep(3)
	return
}