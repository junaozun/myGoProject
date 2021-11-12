package main

import (
	"fmt"
	"math"
)

const abortIndex int8 = math.MaxInt8 / 2

type Context struct {
	handlers []func(c *Context)
	index    int8
}

func (c *Context) Use(f func(c *Context)) {
	c.handlers = append(c.handlers, f)
}

func (c *Context) Next() {
	if c.index < int8(len(c.handlers)) { //这种情况下才剥洋葱
		c.index++
		c.handlers[c.index](c) //剥洋葱，下一个
	}
}

func (c *Context) Run() {
	c.handlers[0](c)
}

func (c *Context) GET(path string, f func(c *Context)) {
	c.handlers = append(c.handlers, f)
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func main() {
	c := &Context{}
	c.Use(Middleware1())
	c.Use(Middleware2())
	c.Use(Middleware3())
	c.GET("/", func(c *Context) {
		fmt.Println("main handler")
	})
	c.Run()
}

func Middleware1() func(c *Context) {
	return func(c *Context) {
		fmt.Println("middle1")
		c.Next()
		fmt.Println("middle1 end")
	}
}

func Middleware2() func(c *Context) {
	return func(c *Context) {
		fmt.Println("middle2")
		c.Abort() // 第三个中间件不会被执行
		c.Next()
		fmt.Println("middle2 end")
	}
}

func Middleware3() func(c *Context) {
	return func(c *Context) {
		fmt.Println("middle3")
		c.Next()
		fmt.Println("middle3 end")
	}
}
