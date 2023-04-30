FROM golang:1.19.8-alpine

RUN apk update && apk --no-cache add git gcc musl-dev && mkdir -p /go/src/github.com/backend && cd /go/src/github.com/backend && go mod init 

WORKDIR /go/src/github.com/backend

ADD . /go/src/github.com/backend/