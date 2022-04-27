#!/bin/bash

DIR="$( cd "$( dirname "$0"  )" && pwd  )"
set -ex
P=$DIR/..
go install ../../cmd/protoc-gen-sxf
protoc --proto_path=$P/example2 --go_out=paths=source_relative:$P/example2 example2.proto
protoc --proto_path=$P/example3 --go_out=paths=source_relative:$P/example3 example3.proto

protoc --proto_path=$P/example3 --proto_path=$P/example2 --proto_path=$DIR/sometype --go_out=paths=source_relative:$DIR/sometype sometype.proto

protoc --proto_path=$P/example3 --proto_path=$P/example2 --proto_path==$DIR/sometype --proto_path=. --go_out=paths=source_relative:. example.proto test.proto
protoc --proto_path=$P/example3 --proto_path=$P/example2 --proto_path==$DIR/sometype --proto_path=. --base_out=plugins=clone,paths=source_relative:. example.proto test.proto
# gen b.pb
PB_OUTPUT=b.pb protoc --proto_path=$P/example3 --proto_path=$P/example2 --proto_path==$DIR/sometype --proto_path=. --base_out=plugins=clone,paths=source_relative:. example.proto test.proto
