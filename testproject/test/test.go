package main

import (
	"fmt"
	"sync"
	"time"
)

type aa struct {
	a int
	b int
}

func calConeKeyProgress(totalValue int64, payArr []int64) []int64 {
	res := calWeightProgress(totalValue, payArr)
	return res[1:]
}

func calWeightProgress(totalValue int64, payArr []int64) []int64 {
	var sum int64
	for _, v := range payArr {
		sum += v
	}
	res := make([]int64, 0, len(payArr))
	var temp int64
	for _, v := range payArr {
		temp += v
		cc := temp * totalValue / sum
		res = append(res, cc)
	}
	return res
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {

			}
		}()
	}
	wg.Wait()
	fmt.Println("88")
}

func aaaa() []int {
	return nil
}
func f1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("recover panic:%v\n", e)
		}
	}()
	// 开启一个goroutine执行任务
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("recover panic222222:%v\n", e)
			}
		}()

		fmt.Println("in goroutine....")
		// 只能触发当前goroutine中的defer
		panic("panic in goroutine")
	}()

	time.Sleep(time.Second)
	fmt.Println("exit")
}
