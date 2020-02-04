FROM golang:1.13
MAINTAINER super "xjc23456789@163.com"
ENV GOPROXY=https://goproxy.cn
COPY ./ /orbbeciot
WORKDIR /orbbeciot/cmd

RUN go mod download && go build -o iot-main

EXPOSE 9898

ENTRYPOINT ["./iot-main"]
