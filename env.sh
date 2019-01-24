#!/bin/sh

export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export PATH=$PATH:$GOPATH/bin

go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc