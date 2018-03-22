FROM golang:1.9.3-alpine3.7

COPY . $GOPATH/src/go-web

WORKDIR $GOPATH/src/go-web/server

RUN go build -o main main.go

ENTRYPOINT ["./main"]

# must use ./main

