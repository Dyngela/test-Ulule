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

ENV APP_PORT=${APP_PORT}
RUN go build -o /ulule
EXPOSE ${APP_PORT}

ENTRYPOINT ["/ulule"]