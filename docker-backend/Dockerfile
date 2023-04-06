FROM golang:1.19.8-alpine

RUN apk update && apk add git
RUN mkdir -p /go/src/github.com/backend

WORKDIR /go/src/github.com/backend

ADD . /go/src/github.com/backend/
RUN go mod init jobmemo


