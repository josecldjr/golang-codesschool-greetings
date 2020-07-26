#!/usr/bin/env bash
# RUN go install
go get github.com/gorilla/mux

go test
go build -o hello *.go