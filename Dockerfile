FROM golang

WORKDIR /go/src/helpy

COPY . /go/src/helpy

RUN go mod download && go build -o ./bin ./cmd/api/main.go

EXPOSE 5579

ENTRYPOINT /go/src/helpy/bin/main