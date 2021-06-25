# Build stage
FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /go/src/helpy
COPY . .
RUN go build -o ./bin/ cmd/api/main.go

# Run stage

FROM alpine:latest
WORKDIR /go/src/helpy
COPY --from=builder /go/src/helpy/bin/main .
COPY ./.env .

EXPOSE 5579
ENTRYPOINT /go/src/helpy/bin/main