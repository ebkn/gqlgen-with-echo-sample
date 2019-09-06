FROM golang:1.11.3-alpine3.8

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN apk update && \
  apk add --no-cache \
    git
