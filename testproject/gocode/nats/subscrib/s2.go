package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	// 连接Nats服务器
	nc, _ := nats.Connect("nats://0.0.0.0:4222")

	// 发布-订阅 模式，异步订阅 test1
	ss, _ := nc.Subscribe("test1", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// 队列 模式，订阅 test2， 队列为queue, test2 发向所有队列，同一队列只有一个能收到消息
	_, _ = nc.QueueSubscribe("test2", "queue", func(msg *nats.Msg) {
		fmt.Printf("Queue a message: %s\n", string(msg.Data))
	})

	// 请求-响应， 响应 test3 消息。
	_, _ = nc.Subscribe("test3", func(m *nats.Msg) {
		fmt.Printf("Reply a message: %s\n", string(m.Data))
		// _ = nc.Publish(m.Reply, []byte("I can help!!"))
		nc.PublishMsg(&nats.Msg{
			Subject: m.Reply,
			Data:    []byte("ni hao a xue"),
		})
	})
	for {

	}
	// 持续发送不需要关闭
	_ = nc.Drain()

	//	关闭连接
	nc.Close()

}
