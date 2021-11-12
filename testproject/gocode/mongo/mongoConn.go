package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func initEngine() {
	var err error

	// 链接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://web.sanguo.bj:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}
