package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-redis/redis"
	"golang.org/x/sync/errgroup"
)

const (
	topic = "test"
)

func TestNewPartitionACKListMQ(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	mq := NewPartitionACKListMQ(client)
	// go mq.Consume(context.Background(), topic, 0, func(msg *Msg) error {
	// 	vs := &MsgData{}
	// 	err := json.Unmarshal(msg.Body, vs)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Printf("consume partiton%d,data:%v", 0, vs)
	// 	return nil
	// })

	eg, ctx1 := errgroup.WithContext(context.Background())
	res := make(chan *MsgData, 10000)
	for i := 0; i < 2; i++ {
		partition := i
		eg.Go(func() error {
			err := mq.Consume(ctx1, topic, partition, func(msg *Msg) error {
				vs := &MsgData{}
				err := json.Unmarshal(msg.Body, vs)
				if err != nil {
					return err
				}
				fmt.Printf("consume partiton%d,data:%v", i, vs)
				res <- vs
				return nil
			})
			if err != nil {
				return err
			}
			return nil
		})
	}
	for {
		select {
		case data := <-res:
			fmt.Println(data)
		}
	}
	err := eg.Wait()
	if err != nil {
		t.Error(err)
	}
}

type CallMsg struct {
	Id            int64  `db:"id" json:"id,omitempty"`
	Status        int8   `db:"status" json:"status,omitempty"`
	FromAddress   string `db:"from_address" json:"from_address,omitempty"`
	BizType       int8   `db:"biz_type" json:"biz_type,omitempty"`
	CreateTime    int64  `db:"create_time" json:"create_time,omitempty"`
	UpdateTime    int64  `db:"update_time" json:"update_time,omitempty"`
	TokenCode     string `db:"token_code" json:"token_code,omitempty"`
	BalanceReal   string `db:"balance_real" json:"balance_real,omitempty"`
	TokenId       int64  `db:"token_id" json:"token_id,omitempty"`
	TxHash        string `db:"tx_hash" json:"tx_hash,omitempty"`
	UnionUserId   string `db:"union_user_id" json:"union_user_id,omitempty"`
	OutWithdrawId string `db:"out_withdraw_id" json:"out_withdraw_id,omitempty"`
	TokenUri      string `db:"token_uri" json:"token_uri,omitempty"`
}

func TestListMQ_Send(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	mq := NewPartitionACKListMQ(client)
	for i := 0; i < 10; i++ {
		msg := CallMsg{
			Id:          1,
			Status:      1,
			FromAddress: "dsndiaiaidjfajf",
			BizType:     1,
		}
		b, err := json.Marshal(msg)
		if err != nil {
			t.Error(err)
			return
		}
		mq.SendMsg(context.Background(), &Msg{
			Topic:     "callback",
			Body:      b,
			Partition: 0,
		})
	}
}
