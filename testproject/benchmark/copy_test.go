package benchmark

import "testing"

func copyFunc() []int {
	temp := make([]int, 0, 10)
	temp = append(temp, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	res := make([]int, 10)
	copy(res, temp)
	return res
}

func valueFunc() []int {
	temp := make([]int, 0, 10)
	temp = append(temp, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	var ls []int
	ls = temp
	return ls
}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyFunc()
	}
}

func BenchmarkValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueFunc()
	}
}
