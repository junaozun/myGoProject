package benchmark

import (
	"fmt"
	"sync"
	"testing"
)

type student struct {
	id int
	score int
}

func BenchmarkCpuA(b *testing.B) {
	str := student{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 100000; i++ {
			str.id++
			str.score++
		}
	}
}

func BenchmarkCpuB(b *testing.B) {
	str := student{}
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				str.id++
			}
		}()
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				str.score++
			}
		}()
		wg.Wait()
	}
	fmt.Println(str)
}