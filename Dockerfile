FROM golang:1.13
MAINTAINER super "xjc23456789@163.com"
ENV GOPROXY=https://goproxy.cn
WORKDIR /orbbeciot
COPY ./ /orbbeciot
RUN ls -al
RUN pwd

RUN go mod download

RUN cd cmd && go build -o main

EXPOSE 9898

ENTRYPOINT ["./cmd/main"]