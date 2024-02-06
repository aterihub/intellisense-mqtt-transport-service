FROM golang:alpine AS build

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /intellisense-mqtt-transport-service
RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o intellisense-mqtt-transport-service

FROM alpine:3.18 as production
RUN apk add --no-cache ca-certificates

COPY --from=build intellisense-mqtt-transport-service /usr/local/bin/
COPY config.json /etc/intellisense-mqtt-transport-service/config.json

ENTRYPOINT [ "intellisense-mqtt-transport-service" ]
CMD ["--config", "/etc/intellisense-mqtt-transport-service/config.json"]