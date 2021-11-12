package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client     = GetMgoCli()
	db         *mongo.Database
	collection *mongo.Collection
	err        error
	ctx        = context.TODO()
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 选择数据库
	db = client.Database("webapi_global_debug")
	// 选择表
	collection = db.Collection("role_extra")

	for i := 0; i < 100; i++ {
		go bulkWrite()
	}
	//insertOne()
	//insertMany()
	//updateOne()
	//findOne()
	//findMany()
	//count()
	//deleteOne()
	//deleteMany()
	//dataAggregate(1)
	time.Sleep(10 * time.Second)
	_ = client.Disconnect(ctx)
}

// 插入一条数据
func insertOne() {

	one, _ := collection.InsertOne(ctx, bson.M{"xuehao": 77777, "name": "你好"})
	one2, _ := collection.InsertOne(ctx, bson.M{"xuehao": 88888, "name": "刘德华"})
	fmt.Println(one.InsertedID, one2.InsertedID)
}

// 批量插入数据
func bulkWrite() {

	type roleDetail struct {
		UID uint64 `bson:"_id"` // id
		SID uint64 `bson:"sid"`
	}

	dis := map[uint64]*roleDetail{
		100111: {
			UID: 10004100062,
			SID: 10089,
		},
	}

	wm := make([]mongo.WriteModel, 0, 100)
	for _, v := range dis {
		filter := bson.D{{"_id", v.UID}}
		temp := mongo.NewReplaceOneModel().
			SetUpsert(true).
			SetFilter(filter).
			SetReplacement(v)
		wm = append(wm, temp)
	}
	ret, err := collection.BulkWrite(ctx, wm)
	if err != nil {
		panic(err)
	}
	var total int64
	if ret != nil {
		total = ret.UpsertedCount + ret.ModifiedCount
	}
	fmt.Println(total)
	_, err = collection.DeleteOne(ctx, bson.M{"_id": 10004100062})
	if err != nil {
		panic(err)
	}
}

// 插入多条数据
func insertMany() {
	documents := []interface{}{
		bson.M{"_id": 1000410054014549, "name": "游龙"},
		bson.M{"_id": 1748310001032513, "name": "侠客"},
	}
	many, _ := collection.InsertMany(ctx, documents)
	fmt.Println(many.InsertedIDs)
}

// 更新数据
// 只要匹配到一个，更新后就返回了，不会把所有匹配的修改
func updateOne() {
	filter := bson.D{{"name", "赵匡胤123"}}
	update := bson.D{{"$set", bson.D{{"name", "赵匡胤新名字123"}}}}
	one, _ := collection.UpdateOne(ctx, filter, update)

	//官方例子 其实还加了个ops参数: options.Update().SetUpsert(true) 该参数的意思 如果没有根据fileter查询到结果，就作为新数据插入其中
	//one, _ := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))

	fmt.Println(one.UpsertedCount, one.MatchedCount, one.ModifiedCount, one.UpsertedID)
}

// 返回找到的第一条记录
func findOne() {
	filter := bson.M{"name": "李世石1"}
	one, _ := collection.FindOne(ctx, filter).DecodeBytes()
	cur, _ := collection.Find(ctx, filter)
	var result []test
	cur.All(ctx, &result)
	fmt.Println(one)
	fmt.Println(result)
}

func findMany() {
	//设置排序
	//opts := options.Find().SetSort(bson.D{{"age", 1}})
	//设置返回结果条数
	//options.Find().SetLimit(10)

	// 不过滤，查询所有数据
	//filter := bson.M{}

	// 过滤
	filter2 := bson.M{"name": "李世石"}

	find, _ := collection.Find(ctx, filter2)
	var results []bson.M
	_ = find.All(ctx, &results)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 统计改集合所有数据总数
func count() {
	documents, _ := collection.CountDocuments(ctx, bson.M{})
	fmt.Println(documents)
}

// 删除一条数据
func deleteOne() {
	one, _ := collection.DeleteOne(ctx, bson.M{"name": "李世石"})
	fmt.Println(one.DeletedCount)
}

// 删除多条数据
func deleteMany() {
	// 过滤条件为空，清空整个集群
	many, _ := collection.DeleteMany(ctx, bson.M{})
	fmt.Println(many.DeletedCount)
}

type test struct {
	Id   int    `bson:"_id"`
	Name string `bson:"name"`
}

type roleInfo struct {
	Id                    int         `bson:"id"`
	OpId                  int         `bson:"opgame_id"`
	Account               string      `bson:"account"`
	Sid                   int         `bson:"sid"`
	ChanId                string      `bson:"chan_id"`
	Name                  string      `bson:"name"`
	Lv                    int         `bson:"lv"`
	VipLv                 int         `bson:"vipLv"`
	CreateTime            int         `bson:"createTime"`
	LastLogin             int         `bson:"lastLogin"`
	LastLogout            int         `bson:"lastLogout"`
	VipExp                int         `bson:"vipExp"`
	Diamond               int         `bson:"diamond"`
	Power                 int         `bson:"power"`
	MaxPower              int         `bson:"max_power"`
	Avatar                int         `bson:"avatar"`
	Gold                  int         `bson:"gold"`
	LastRechargeTime      int         `bson:"lastRechargeTime"`
	TodayRechargeMoney    int         `bson:"todayRechargeMoney"`
	TotalRechargeDiamond  int         `bson:"totalRechargeDiamond"`
	TotalRechargeMoney    int         `bson:"totalRechargeMoney"`
	UpTime                int         `bson:"up_time"`
	ActiveScore           int         `bson:"active_score"`
	ActiveScoreUpdateTime int         `bson:"active_score_update_time"`
	Extras                []roleExtra `bson:"extras"`
}

type roleExtra struct {
	Id   int    `bson:"_id"`
	Name string `bson:"name"`
}

/*
[
  {
    "_id": 1748310001032513,
    "opgame_id": 1000,
    "account": "0060015_15761568671000006691",
    "sid": 1748310001,
    "channel_id": "",
    "name": "",
    "lv": 0,
    "vipLv": 0,
    "createTime": 0,
    "lastLogin": 0,
    "lastLogout": 1633259529,
    "vipExp": 7860,
    "diamond": 41212,
    "power": 47905621,
    "max_power": 19680008,
    "avatar": 1058,
    "gold": 1029627587,
    "lastRechargeTime": 1631123932,
    "todayRechargeMoney": 0,
    "totalRechargeDiamond": 680,
    "totalRechargeMoney": 786,
    "up_time": 1633259529,
    "active_score": 0,
    "active_score_update_time": 1633259529,
    "ids": [
      {
        "_id": 1748310001032513,
        "name": "侠客"
      }
    ]
  }
]
*/
// 联表查询
func dataAggregate(id int) {
	pipe := genPipeline(id)
	opts := options.Aggregate().SetMaxTime(15 * time.Second)
	cursor, err := collection.Aggregate(context.TODO(), pipe, opts)
	if err != nil {
		log.Fatal(err)
	}

	//打印文档内容
	//var results []bson.M
	//if err = cursor.All(context.TODO(), &results); err != nil {
	//	log.Fatal(err)
	//}
	//for _, result := range results {
	//	fmt.Println(result)
	//}

	// 解析到结构体中
	var roleInfos []roleInfo
	cursor.All(ctx, &roleInfos)
	fmt.Println(roleInfos)

	for _, v := range roleInfos {
		for _, vv := range v.Extras {
			fmt.Println(vv.Name)
		}
	}
}

/*
$lookup：     左联表字段
from：        需要连接的集合（tableOther）
localField：  在输入文档（collectionMain）中的查找字段
foreignField：需要在from集合中查找的字段
as：          输出的字段名字，可任你更改
*/
func genPipeline(id int) mongo.Pipeline {
	matchStat := bson.D{
		{
			"$lookup", bson.D{
				{"from", "role_extra"},
				{"localField", "_id"},
				{"foreignField", "_id"},
				{"as", "extras"},
			},
		},
	}
	filter := bson.D{
		{
			"$match", bson.D{
				{"_id", 1748310001032513},
			},
		},
	}
	return mongo.Pipeline{matchStat, filter}
}
