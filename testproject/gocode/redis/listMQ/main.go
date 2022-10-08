package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Msg struct {
	Topic     string // 消息的主题
	Body      []byte // 消息的Body
	Partition int    // 分区号
}

type PartitionACKListMQ struct {
	client *redis.Client // Redis客户端
}

func NewPartitionACKListMQ(client *redis.Client) *PartitionACKListMQ {
	return &PartitionACKListMQ{client: client}
}

func (q *PartitionACKListMQ) SendMsg(ctx context.Context, msg *Msg) error {
	return q.client.LPush(partitionTopic(msg.Topic, msg.Partition), msg.Body).Err()
}

func partitionTopic(topic string, partition int) string {
	return fmt.Sprintf("%s:%d", topic, partition)
}

type Handler func(msg *Msg) error

func (q *PartitionACKListMQ) Consume(ctx context.Context, topic string, partition int, h Handler) error {
	for {
		// 获取消息
		body, err := q.client.LIndex(partitionTopic(topic, partition), -1).Bytes()
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}
		// 没有消息了，休眠一会
		if errors.Is(err, redis.Nil) {
			time.Sleep(time.Second)
			continue
		}
		// 处理消息
		err = h(&Msg{
			Topic:     topic,
			Body:      body,
			Partition: partition,
		})
		if err != nil {
			continue
		}
		// 如果处理成功，删除消息
		if err := q.client.RPop(partitionTopic(topic, partition)).Err(); err != nil {
			return err
		}
	}
}
