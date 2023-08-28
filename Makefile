.PHONY: run build test lint compose

# Service port
PORT ?= 10100

run:
	HTTP_ADDR=:$(PORT) go run "cmd/promotions_service/main.go"

build:
	go build -o ".bin/promotions_service" "cmd/promotions_service/main.go"

test:
	go test -race ./...

lint:
	golangci-lint run -c .golangci.yml -v ./...

compose:
	docker-compose up  --remove-orphans --build
