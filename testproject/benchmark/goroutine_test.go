package benchmark

import (
	"sync"
	"testing"
)

func BenchmarkGoroutineOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		initGoroutine1()
	}
}

func initGoroutine1() {
	for i := 0; i < 20; i++ {
		a := 0
		for i := 0; i < 10000000; i++ {
			a += i
		}
	}
}

func BenchmarkGoroutineMut(b *testing.B) {
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		initGoroutine2(wg)
	}
}

func initGoroutine2(wg sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			cd := 0
			for i := 0; i < 10000000; i++ {
				cd += i
			}
		}(i)
	}
	wg.Wait()
}
