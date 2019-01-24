#!/bin/sh

if [ "x$GOPATH" = "x" ]; then
    echo "ERROR : GOPATH is not set"
    exit 1
fi

go get github.com/golang/protobuf/protoc-gen-go
go get google.golang.org/grpc

protoc --plugin=grpc=$GOPATH/bin/protoc-gen-go -I=$GOPATH/src/github.com/Thommil/grpc-web service/service.proto --go_out=plugins=grpc:$GOPATH/src/github.com/Thommil/grpc-web

cd $GOPATH/src/github.com/Thommil/grpc-web/server
go build -o $GOPATH/bin/server main.go

cd $GOPATH/src/github.com/Thommil/grpc-web/client
go build -o $GOPATH/bin/client main.go