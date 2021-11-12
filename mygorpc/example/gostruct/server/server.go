package main

import (
	"mygorpc"
	"mygorpc/example/gostruct/way"
	"time"
)

func main() {
	opts := []mygorpc.ServerOption{
		mygorpc.WithAddress("127.0.0.1:8000"),
		mygorpc.WithNetwork("tcp"),
		mygorpc.WithSerializationType("msgpack"),
		mygorpc.WithTimeout(time.Millisecond * 2000),
	}
	s := mygorpc.NewServer(opts...)
	if err := s.RegisterService("/helloworld.Greeter", new(way.Service)); err != nil {
		panic(err)
	}
	s.Serve()
}
