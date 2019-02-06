FROM golang:1.11.3-alpine3.8

ENV GO111MODULE on

WORKDIR /go/src/app

COPY . .

RUN apk add --no-cache \
  git \
  gcc
