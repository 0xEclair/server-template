FROM golang:1.18-alpine

ADD . /server-template

WORKDIR /server-template

ENTRYPOINT ["sh", "-c", "go run main.go"]