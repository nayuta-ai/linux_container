# FROM golang:1.19.1-alpine
FROM --platform=linux/amd64 ubuntu:20.04

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
ENV GO_VERSION=1.19

RUN apt-get update
RUN apt-get install -y wget git gcc xz-utils net-tools iproute2

RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR /usr/src

RUN wget "https://github.com/teddyking/netsetgo/releases/download/0.0.1/netsetgo" \
    && mv netsetgo /usr/local/bin/ \
    && chown root:root /usr/local/bin/netsetgo \
    && chmod 4755 /usr/local/bin/netsetgo