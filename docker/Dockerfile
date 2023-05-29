# FROM golang:1.19.1-alpine
FROM ubuntu:20.04

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
ENV GO_VERSION=1.19

RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-arm64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-arm64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-arm64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR /usr/src