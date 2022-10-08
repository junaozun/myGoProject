package main

import (
	"log"
	"strings"

	"github.com/nats-io/nats.go"
)

func main() {
	// 1.配置参数
	// 链接地址
	// var urls = flag.String("s", "nats://10.18.98.163:4222,nats://10.18.98.164:4222,nats://10.18.98.165:4222", "nats urls")
	// // var tls = flag.Bool("tls",false,"是否使用安全传输")
	// log.SetFlags(log.Ldate)
	// flag.Parse()
	// args := flag.Args()
	// if len(args) < 1 {
	// }

	// 配置可选项
	opts := nats.GetDefaultOptions()
	opts.Servers = strings.Split("0.0.0.0:4222", ",")
	for k, s := range opts.Servers {
		opts.Servers[k] = strings.Trim(s, " ")
	}

	// 链接到nats
	nc, err := opts.Connect()
	if err != nil {
		log.Fatalf("Cago envn not connect:%v\n", err)
	}

	// 发布主题
	subject := "su"
	data := "我是程序员"
	nc.Publish(subject, []byte(data))
	nc.Flush()
}
