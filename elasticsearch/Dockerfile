FROM golang:1.18.0-buster as build

RUN apt-get update

COPY . /elasticsearch
RUN cd /elasticsearch && go build -o service main.go

FROM debian:buster-slim

RUN mkdir -p /opt/elasticsearch
COPY --from=build /elasticsearch/service /opt/elasticsearch/service

WORKDIR /opt/elasticsearch
CMD ["./service"]