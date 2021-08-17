FROM golang:latest

ENV GOPATH=/go
WORKDIR /go/src

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@latest