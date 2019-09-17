FROM golang:1.13-alpine3.10 AS builder-base

RUN apk update && \
    apk add protobuf && \
    apk add git && \
    go get -u -v github.com/golang/protobuf/protoc-gen-go && \
    go get google.golang.org/grpc

FROM builder-base as builder 
        
WORKDIR /go/src/github.com/calumball/ori/
COPY . .

RUN protoc --go_out=plugins=grpc:. counter/proto/counter.proto
