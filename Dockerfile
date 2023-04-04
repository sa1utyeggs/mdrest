FROM golang:1.17-alpine

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

# 将当前目录下的文件复制到 go/src/...
COPY ./ /go/src/github.com/sa1utyeggs/mdrest
WORKDIR /go/src/github.com/sa1utyeggs/mdrest/mdrest
# 生成可执行文件，文件在 GOBIN 目录下
RUN go install -ldflags '-s -w'

FROM alpine:3.14
# 将 GOBIN 目录下的可执行文件 "mdrest" 放入 /app 中
COPY --from=0 /go/bin/mdrest /app/mdrest
COPY --from=0 /go/src/github.com/sa1utyeggs/mdrest/mdrest/config.json /app/
CMD  ["/app/mdrest", "--help"]