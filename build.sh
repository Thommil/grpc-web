#!/bin/sh

protoc --plugin=grpc=$GOPATH/bin/protoc-gen-go -I=$GOPATH service/service.proto --go_out=plugins=grpc:$GOPATH/service

go build main.go