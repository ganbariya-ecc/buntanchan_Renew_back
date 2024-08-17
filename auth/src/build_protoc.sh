#!/bin/bash
# apt update
apt install protobuf-compiler

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

cd ./sdks/proto
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. authsdk.proto