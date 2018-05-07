FROM golang:1.9.3-alpine3.7

# ENV HTTP_PROXY http://95.163.202.160:1080
# ENV HTTPS_PROXY http://95.163.202.160:1080

RUN apk add --no-cache git curl \
    && curl https://glide.sh/get | sh

# RUN go get -u google.golang.org/grpc

COPY . $GOPATH/src/github.com/sundayfun/go-web

WORKDIR $GOPATH/src/github.com/sundayfun/go-web

# RUN glide up
# RUN glide install -v

WORKDIR $GOPATH/src/github.com/sundayfun/go-web/main-server

RUN go build -o web-crawler web-crawler.go 

ENTRYPOINT ["./web-crawler"]

# must use ./main

