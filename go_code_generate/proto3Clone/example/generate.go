//go:generate go install ../../cmd/protoc-gen-sxf
//go:generate protoc --go_out=paths=source_relative:. example.proto
//go:generate protoc --proto_path=. --sxf_out=plugins=clone,paths=source_relative:. example.proto
package example
