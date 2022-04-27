package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	vms := []int32{1,2,3,4,5,6,7,8,9}
	patchInfo := make([]int32,0)
	for _,v := range vms {
		wg.Add(1)
		go func(a int32) {
			defer wg.Done()
			lock.Lock()
			patchInfo = append(patchInfo, a+100)
			lock.Unlock()
		}(v)
	}
	wg.Wait()
	fmt.Println(patchInfo)
}


