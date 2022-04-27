package main

import (
	generator "go_gen/generate"
	proto3clone "go_gen/proto3Clone"
)

func main() {
	generator.Main(&proto3clone.Plugin{})
}
