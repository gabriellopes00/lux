FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /go/src/helpy
COPY . .
RUN apk add make
RUN make install && make build

FROM alpine:3.14
WORKDIR /go/src/helpy
COPY --from=builder /go/src/helpy/bin/app .
COPY .env .

EXPOSE 5590
ENTRYPOINT /go/src/helpy/app