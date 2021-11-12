package main

import (
	"context"
	"fmt"
	"mygorpc"
	"time"

	"github.com/lubanproj/gorpc/examples/helloworld2/helloworld"
)

type greeterService struct{}

func (g *greeterService) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Println("recv Msg : ", req.Msg)
	rsp := &helloworld.HelloReply{
		Msg: req.Msg + " world",
	}
	return rsp, nil
}

func main() {
	opts := []mygorpc.ServerOption{
		mygorpc.WithAddress("127.0.0.1:8000"),
		mygorpc.WithNetwork("tcp"),
		mygorpc.WithProtocol("proto"),
		mygorpc.WithTimeout(time.Millisecond * 2000),
	}
	s := mygorpc.NewServer(opts...)
	helloworld.RegisterService(s, &greeterService{})
	s.Serve()
}
