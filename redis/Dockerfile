FROM golang:1.18.0-buster as build

RUN apt-get update

COPY . /redis
RUN cd /redis && go build -o service main.go

FROM debian:buster-slim

RUN mkdir -p /opt/redis
COPY --from=build /redis/service /opt/redis/service

WORKDIR /opt/redis
CMD ["./service"]