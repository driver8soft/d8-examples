#!/bin/bash
export GOPATH=$(go env GOPATH)

export PATH=$PATH:$GOPATH/bin

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    hello/hello.proto
