FROM golang:1.15-rc-alpine3.12

RUN apk add git

COPY go.mod /ray-tracing/
COPY go.sum /ray-tracing/

WORKDIR /ray-tracing

RUN go mod download

COPY . /ray-tracing

ENV CGO_ENABLED 0
ENV NUMBER_OF_THREADS 4

ENTRYPOINT ["./entrypoint.sh"]
