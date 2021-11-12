package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
红包算法要求：

		1、红包序列是随机的，序列元素之间差异性可以控制
		2、大小分布可以预期
		3、尽可能降低性能损耗
*/

/*
	与wechat红包算法一致：二倍均值算法
	思路：1、每次随机金额的平均值是基本相等
 		 2、剩余金额平均数的2倍作为随机最大数
*/

const min = 1 //最小0.01RMB

func main() {
	// 示列，将100块分成10份
	fmt.Println(CalRedPkg(5, 20))
}

// 计算红包
func CalRedPkg(count, amount int64) []float64 {
	amount = amount * 100 // 将元变成分
	red := make([]float64, 0, count)
	remain := amount
	for i := int64(0); i < count; i++ {
		x := DoubleAverage(count-i, remain)
		remain -= x
		red = append(red, float64(x)/float64(100))
	}
	return red
}

// 二部均值算法
func DoubleAverage(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	//计算出最大可用金额
	max := amount - min*count
	//计算出最大可用平均值
	avg := max / count
	//二倍均值基础再加上最小金额，防止出现0值
	avg2 := 2*avg + min
	//随机红包金额序列元素，把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg2) + min
	return x
}
