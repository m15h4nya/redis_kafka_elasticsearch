FROM golang:1.18.0-buster as build

RUN apt-get update && apt-get install gcc

COPY . /kafka
RUN cd /kafka && go build -o service main.go

FROM debian:buster-slim

RUN mkdir -p /opt/kafka
COPY --from=build /kafka/service /opt/kafka/service

WORKDIR /opt/kafka
CMD ["./service"]