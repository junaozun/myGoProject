package main

import (
	"fmt"
	"sync"
)

func main() {
	var syncMap sync.Map
	v, ok := syncMap.Load(100)
	fmt.Println(v, ok)
	syncMap.Store(100, "nihao")
	v2, ok2 := syncMap.Load(100)
	fmt.Println(v2, ok2)
}
