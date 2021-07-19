FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /go/src/lux
COPY . .
RUN apk update && apk upgrade && apk add make
RUN make install && make build

FROM alpine:3.14
WORKDIR /go/src/lux
COPY --from=builder /go/src/lux/bin/app .
COPY .env .

EXPOSE 5590
ENTRYPOINT /go/src/lux/app