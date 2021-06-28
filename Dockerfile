FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /go/src/helpy
COPY . .
RUN go get ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/app cmd/api/main.go
RUN ls -a

FROM alpine
WORKDIR /go/src/helpy
COPY --from=builder /go/src/helpy/bin/app .
COPY ./.env .

EXPOSE 5590
ENTRYPOINT /go/src/helpy/app