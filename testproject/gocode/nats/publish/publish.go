package main

import (
	"flag"
	"github.com/nats-io/nats.go"
	"log"
	"strings"
)

func main() {
	// 1.配置参数
	// 链接地址
	var urls = flag.String("s","nats://0.0.0.0:4222","nats urls")
	//var tls = flag.Bool("tls",false,"是否使用安全传输")
	log.SetFlags(log.Ldate)
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
	}

	// 配置可选项
	opts := nats.GetDefaultOptions()
	opts.Servers = strings.Split(*urls,",")
	for k,s := range opts.Servers {
		opts.Servers[k] = strings.Trim(s," ")
	}

	// 链接到nats
	nc,err := opts.Connect()
	if err != nil {
		log.Fatalf("Cago envn not connect:%v\n",err)
	}

	// 发布主题
	subject := args[0]
	data := args[1]
	nc.Publish(subject,[]byte(data))
    nc.Flush()
}