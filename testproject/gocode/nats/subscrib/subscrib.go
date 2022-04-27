package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/nats-io/nats.go"
)

func usage() {
	log.Fatalf("Usage:nats-sub [-s server] [--tls]")
}

func main() {
	// 1.配置参数
	// 链接地址
	var urls = flag.String("s", "nats://10.18.98.163:4222,nats://10.18.98.164:4222,nats://10.18.98.165:4222", "nats urls")
	// var tls = flag.Bool("tls",false,"是否使用安全传输")
	log.SetFlags(log.Ldate)
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	// 配置可选项
	opts := nats.GetDefaultOptions()
	opts.Servers = strings.Split(*urls, ",")
	for k, s := range opts.Servers {
		opts.Servers[k] = strings.Trim(s, " ")
	}

	// 链接到nats
	nc, err := opts.Connect()
	if err != nil {
		log.Fatalf("Can not connect:%v\n", err)
	}

	// 订阅主题
	subject := args[0]
	i := 0
	sub, _ := nc.Subscribe(subject, func(msg *nats.Msg) {
		i++
		fmt.Printf("[#%d] Received on [%s]:%s\n", i, msg.Subject, string(msg.Data))
	})

	// 取消订阅
	// 订阅某个主题，收到三个3个这个主题后，自动取消订阅了
	sub.AutoUnsubscribe(3)
	log.Printf("")
	// 等待接受主题消息
	runtime.Goexit()
}
