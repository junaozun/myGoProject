package options

import (
	"fmt"
)

type server struct {
	opt     *options
	servers map[int32]int32
}

type options struct {
	bucketCount uint64
	bucketMask  uint64
}

type Opt func(options *options)

func NewServer(opts ...Opt) *server {
	o := &options{
		bucketCount: 1,
		bucketMask:  1,
	}
	for _, each := range opts {
		each(o)
	}
	return &server{
		opt:     o,
		servers: map[int32]int32{},
	}
}

func (s *server) Println() {
	fmt.Println("server", s.servers)
}

func SetBucketCount(count uint64) Opt {
	return func(opt *options) {
		opt.bucketCount = count
	}
}

func SetBucketMask(mask uint64) Opt {
	return func(opt *options) {
		opt.bucketMask = mask
	}
}

func main() {
	server := NewServer(SetBucketCount(10), SetBucketMask(333))
	server.Println()
}
