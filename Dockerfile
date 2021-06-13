FROM golang

COPY . /go/src/helpy

RUN go mod download
RUN go install helpy

ENTRYPOINT /go/bin/helpy

EXPOSE 5579