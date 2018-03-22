FROM golang:1.9.3-alpine3.7

COPY . $GOPATH/src/database-work

WORKDIR $GOPATH/src/database-work/server

RUN go build -o main main.go

ENTRYPOINT ["./main"]

# must use ./main

