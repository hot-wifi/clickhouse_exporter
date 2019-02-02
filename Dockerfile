FROM golang:1.11.5-alpine AS build-env

MAINTAINER Fadi Hadzh <f.hadzh@hot-wifi.ru>

WORKDIR /go/src/github.com/hot-wifi/clickhouse_exporter

COPY . .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go get -d -v ./...
RUN go build -v .

FROM alpine:3.8

RUN apk --no-cache add tzdata ca-certificates

WORKDIR /app/clickhouse_exporter

COPY --from=build-env /go/src/github.com/hot-wifi/clickhouse_exporter/clickhouse_exporter /bin/clickhouse_exporter

ENTRYPOINT [ "clickhouse_exporter" ]
