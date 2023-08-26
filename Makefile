.PHONY: run build test lint docker-compose

# Service port
PORT ?= 10100

run:
	HTTP_ADDR=:$(PORT) go run "cmd/promotions-service/main.go"

build:
	go build -o ".bin/promotions-service" "cmd/promotions-service/main.go"

test:
	go test -race ./...

lint:
	golangci-lint run -c .golangci.yml -v ./...

docker-compose:
	docker-compose up  --remove-orphans --build