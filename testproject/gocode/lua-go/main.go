package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
)

func main() {
	l := lua.NewState()
	defer l.Close()
	if err := l.DoString(`print("Hello World")`); err != nil {
		panic(err)
	}
}

// 编译 lua 代码字段
func CompileString(source string) (*lua.FunctionProto, error) {
	reader := strings.NewReader(source)
	chunk, err := parse.Parse(reader, source)
	if err != nil {
		return nil, err
	}
	proto, err := lua.Compile(chunk, source)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

// 编译 lua 代码文件
func CompileFile(filePath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

func BenchmarkRunWithoutPreCompiling(b *testing.B) {
	l := lua.NewState()
	for i := 0; i < b.N; i++ {
		_ = l.DoString(`a = 1 + 1`)
	}
	l.Close()
}

func BenchmarkRunWithPreCompiling(b *testing.B) {
	l := lua.NewState()
	proto, _ := CompileString(`a = 1 + 1`)
	lfunc := l.NewFunctionFromProto(proto)
	for i := 0; i < b.N; i++ {
		l.Push(lfunc)
		_ = l.PCall(0, lua.MultRet, nil)
	}
	l.Close()
}

func BenchmarkRunWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := lua.NewState()
		_ = l.DoString(`a = 1 + 1`)
		l.Close()
	}
}

func BenchmarkRunWithPool(b *testing.B) {
	pool := newVMPool(nil, 100)
	for i := 0; i < b.N; i++ {
		l := pool.get()
		_ = l.DoString(`a = 1 + 1`)
		pool.put(l)
	}
}
