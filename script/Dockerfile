FROM golang:1.18.3-alpine3.16 AS build
RUN apk update && \
    apk add --no-cache \
    netcat-openbsd \
    iproute2 \
    nmap \
    curl

COPY . /project
WORKDIR /project
RUN go mod download

ARG BUILD_ARGS
ENV BUILD_ARGS=${BUILD_ARGS}
ARG GO_MAIN
ENV GO_MAIN=${GO_MAIN}