# STEP 1 - Build
FROM golang:1.21.0-alpine as build

WORKDIR /app

ADD . /app/

RUN go mod download
RUN go build -o geolocation-api ./cmd/api
RUN go build -o geolocation-parser ./cmd/parser

CMD ["/app/geolocation-api"]

