package benchmark

import (
	"testing"
)

func BenchmarkFindInArray(b *testing.B) {
	a := make([]int32, 16)
	for i := 0; i < len(a); i++ {
		a[i] = int32(i + 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findInArray(a, 12)
	}
}

func findInArray(a []int32, x int32) int32 {
	for _, v := range a {
		if v == x {
			return v
		}
	}
	return -1
}

func BenchmarkFindInMap(b *testing.B) {
	a := make(map[int32]int32, 16)
	for i := 0; i < len(a); i++ {
		a[int32(i+1)] = int32(i + 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findInMap(a, 12)
	}
}

func findInMap(a map[int32]int32, x int32) int32 {
	if v, ok := a[x]; ok {
		return v
	}
	return -1
}
