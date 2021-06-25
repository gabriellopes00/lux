test:
	go test -v -cover ./...

server:
	go run cmd/api/main.go

.PHONY: test server