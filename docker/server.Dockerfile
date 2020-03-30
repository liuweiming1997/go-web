FROM golang:1.9.3-alpine3.7

# test proxy or not
# RUN go get -u google.golang.org/grpc

COPY . $GOPATH/src/github.com/sundayfun/go-web

WORKDIR $GOPATH/src/github.com/sundayfun/go-web

# WORKDIR $GOPATH/src/github.com/sundayfun/go-web/main-server

# RUN go build -o web-crawler web-crawler.go

COPY ./shell/setup.sh /usr/local/bin

ENTRYPOINT ["setup.sh"]

# must use ./main
