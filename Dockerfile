FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

ENV GO111MODULE="on" \
    GOPROXY="https://goproxy.cn,direct"

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8888

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
