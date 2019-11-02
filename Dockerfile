FROM golang:1.13-alpine as builder
MAINTAINER jerry hzde0128@live.cn
ADD . /go/src/github.com/hzde0128/gimg
WORKDIR /go/src/github.com/hzde0128/gimg
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk update && apk add gcc musl-dev imagemagick6-dev && go build -v build/main.go

FROM alpine:3.10
MAINTAINER jerry hzde0128@live.cn
WORKDIR /app
ADD build/conf conf
COPY --from=builder /go/src/github.com/hzde0128/gimg/main gimg
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk update && apk add tzdata && apk add imagemagick6-dev && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ENV PATH /app:$PATH
VOLUME /data
EXPOSE 4869
CMD ["gimg", "-config", "conf/config.ini"]