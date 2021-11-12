package benchmark

import (
	"strings"
	"testing"
)

func BenchmarkStringBuildAddGrow(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		stringBuildAddGrow()
	}
}

func stringBuildAddGrow()  {
	var b strings.Builder
	b.Grow(43)
	b.WriteString("昵称")
	b.WriteString(":")
	b.WriteString("志强1224")
	b.WriteString("\n")
	b.WriteString("联系方式QQ")
	b.WriteString(":")
	b.WriteString("354662600")
	b.WriteString("\n")
}

/*
➜  benchmark go test -bench=. stringBuild_test.go -benchmem
goos: darwin
goarch: amd64
BenchmarkStringBuildAddGrow-8           35614303                34.2 ns/op            48 B/op          1 allocs/op
PASS
ok      command-line-arguments  2.376s

*/