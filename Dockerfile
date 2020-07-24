FROM golang:1.13.12-alpine3.12

WORKDIR /go/src

COPY . /go/src
RUN apk add git
RUN apk add gcc
RUN apk add bash
RUN go get github.com/gorilla/mux
# RUN go get ./