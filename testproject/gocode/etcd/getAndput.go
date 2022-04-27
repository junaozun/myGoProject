package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var cli *clientv3.Client

func init() {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.18.98.163:2379", "10.18.98.164:2379", "10.18.98.165:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
}

func main() {

	// putEtcd()
	// getEtcd()
	putEtcdWithLease()

	cli.Close()
}

func putEtcd() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, "/suxf/q1mi", "dsb")
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	cancel()
}

// 设置具有过期时间的key
func putEtcdWithLease() {
	lease := clientv3.NewLease(cli)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// 过期时间10秒
	grantResp, err := lease.Grant(ctx, 10)
	if err != nil {
		fmt.Printf("grantResp failed, err:%v\n", err)
		return
	}
	_, err = cli.Put(ctx, "/suxf/nihao", "这你是的益达", clientv3.WithLease(grantResp.ID))
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	cancel()
}

// func getEtcd() {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	resp, err := cli.Get(ctx, "/game/sanguo/dev/service/state", clientv3.WithPrefix())
// 	cancel()
// 	if err != nil {
// 		fmt.Printf("get from etcd failed, err:%v\n", err)
// 		return
// 	}
// 	for _, ev := range resp.Kvs {
// 		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
// 	}
// }

func watch() {
	// watch key:q1mi change
	rch := cli.Watch(context.Background(), "sxf") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
