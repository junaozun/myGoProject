package benchmark

import "testing"

func normalCopy(arr []int)[]int {
	temp := make([]int, len(arr))
	copy(temp,arr)
	return temp
}

func lineCopy(arr []int)[]int {
	return append(arr[:0:0],arr...)
}

func appendCopy(arr []int)  []int{
	temp := make([]int,len(arr))
	for i,v := range arr {
		temp[i] = v
	}
	return temp
}


func genArray() []int{
	arr := make([]int,50)
	for i,_ := range arr {
		arr[i] = 1000
	}
	return arr
}

func BenchmarkNormal(b *testing.B) {
	arr := genArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		normalCopy(arr)
	}
}

func BenchmarkLine(b *testing.B) {
	arr := genArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lineCopy(arr)
	}
}

func BenchmarkAppend(b *testing.B) {
	arr := genArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		appendCopy(arr)
	}
}
