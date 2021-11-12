package main

import (
	"fmt"
	"strings"
)

var (
	saveMapKey = "cross:gvg:%d:maps:%s"
)

func main() {
	f := EfficientSplicing("ssfdsf", "dsafdsaf")
	fmt.Println(f)
	d := sf(saveMapKey, 10, "34")
	fmt.Println(d)
}

// 高效字符串拼接
func EfficientSplicing(option ...string) string {
	var b strings.Builder
	b.Grow(len(option))
	for _, s := range option {
		b.WriteString(s)
	}
	return b.String()
}

func sf(key string, option ...interface{}) string {
	return fmt.Sprintf(key, option...)
}
