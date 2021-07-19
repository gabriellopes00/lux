pgup:
	@docker run -d \
	--name lux_db \
	-e POSTGRES_PASSWORD=psql \
	-e POSTGRES_DB=lux \
	-v luxdb:/var/lib/postgresql/data/ \
	-p 5432:5432 \
  postgres

pgdown:
	@docker rm -f lux_db

composeup:
	@docker-compose up --build -d

composedown:
	@docker-compose down

test:
	@go test -v -cover ./...

install:
	@go get ./...

server:
	@go run cmd/api/main.go

build:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/app cmd/api/main.go

.PHONY: pgdown pgup test install server build composeup composedown