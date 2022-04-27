package benchmark

import (
	"testing"
)

const (
	maxRow int = 10000
	maxCol int = 10000
)

func BenchmarkRow(b *testing.B) {
	m := initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i]); j++ {
				_ = m[i][j]
			}
		}
	}
}

func BenchmarkCol(b *testing.B) {
	m := initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < maxCol; i++ {
			for j := 0; j < maxRow; j++ {
				_ = m[j][i]
			}
		}
	}
}

func initData() [][]int {
	a := make([][]int, maxRow) // 二维切片，3行
	for i := range a {
		a[i] = make([]int, maxCol) // 每一行4列
	}
	return a
}
