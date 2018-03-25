FROM golang:1.9.3-alpine3.7

COPY . $GOPATH/src/github.com/sundayfun/go-web

WORKDIR $GOPATH/src/github.com/sundayfun/go-web/server

RUN go build -o main main.go

ENTRYPOINT ["./main"]

# must use ./main

