FROM golang:1.9.3-alpine3.7

RUN apk add --no-cache git curl \
    && curl https://glide.sh/get | sh

# test proxy or not
# RUN go get -u google.golang.org/grpc

COPY . $GOPATH/src/github.com/sundayfun/go-web

WORKDIR $GOPATH/src/github.com/sundayfun/go-web

COPY ./docker/mirrors.yaml /root/.glide/mirrors.yaml
RUN glide mirror set https://golang.org/x/crypto https://github.com/golang/crypto --vcs git \
    && glide mirror set https://golang.org/x/net https://github.com/golang/net --vcs git \
    && glide mirror set https://golang.org/x/sys https://github.com/golang/sys --vcs git \
    && glide mirror set https://golang.org/x/text https://github.com/golang/text --vcs git \
    && glide mirror set https://google.golang.org/grpc https://github.com/grpc/grpc-go --vcs git \
    && glide mirror set https://google.golang.org/genproto https://github.com/google/go-genproto --vcs git \
    && glide mirror set https://golang.org/x/image https://github.com/golang/image --vcs git 
RUN glide up
RUN glide install -v

WORKDIR $GOPATH/src/github.com/sundayfun/go-web/main-server

RUN go build -o web-crawler web-crawler.go 

ENTRYPOINT ["./web-crawler"]

# must use ./main

