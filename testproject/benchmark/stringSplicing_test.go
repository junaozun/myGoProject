package benchmark

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)

	}
	return s
}

func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

func preByteConcat(n int, str string) string {
	buf := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

func builderConcatGrow(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func joinString(n int, str string) string {
	var str2 string
	for i := 0; i < n; i++ {
		str2 = strings.Join([]string{str}, str)
	}
	return str2
}

func benchmark(b *testing.B, f func(int, string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

func BenchmarkPlusConcat(b *testing.B) {
	benchmark(b, plusConcat)
}

func BenchmarkSprintfConcat(b *testing.B) {
	benchmark(b, sprintfConcat)
}

func BenchmarkBuilderConcat(b *testing.B) {
	benchmark(b, builderConcat)
}

func BenchmarkBufferConcat(b *testing.B) {
	benchmark(b, bufferConcat)
}

func BenchmarkByteConcat(b *testing.B) {
	benchmark(b, byteConcat)
}

func BenchmarkPreByteConcat(b *testing.B) {
	benchmark(b, preByteConcat)
}

func BenchmarkBuilderConcatGrow(b *testing.B) {
	benchmark(b, builderConcatGrow)
}

func BenchmarkJoinString(b *testing.B) {
	benchmark(b, joinString)
}

/*
➜  benchmark go test -bench=. stringBuild_test.go -benchmem
goos: darwin
goarch: amd64
pkg: testproject/benchmark
BenchmarkPlusConcat-8                 25          40639958 ns/op        530998876 B/op     10034 allocs/op
BenchmarkSprintfConcat-8              15          71384822 ns/op        833860826 B/op     37443 allocs/op
BenchmarkBuilderConcat-8           14061             84918 ns/op          522226 B/op         23 allocs/op
BenchmarkBufferConcat-8            12714             94222 ns/op          423538 B/op         13 allocs/op
BenchmarkByteConcat-8              13602             88141 ns/op          628723 B/op         24 allocs/op
BenchmarkPreByteConcat-8           25610             45025 ns/op          212992 B/op          2 allocs/op
BenchmarkBuilderConcatGrow-8       26414             45633 ns/op          106496 B/op          1 allocs/op
PASS
ok      testproject/benchmark   12.455s
*/
