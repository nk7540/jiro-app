FROM golang:1.15

ENV GO111MODULE=auto

RUN apt update
RUN apt install -y protobuf-compiler \
    vim

WORKDIR /go/artics-api

ADD . /go/artics-api

RUN go mod tidy
