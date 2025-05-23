.PHONY: run migrate-up build docker-up docker-down lint test clean

APP_NAME = todo
DOCKER_COMPOSE_FILE = docker-compose.yaml

run:
	go run ./cmd/todo/main.go

migrate-up:
	go run ./cmd/todo/main.go migrate-up

build:
	go build -o bin/$(APP_NAME) ./cmd/todo

docker-up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

docker-down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

lint:
	golangci-lint run --fix ./...
test:
	go test -v ./...

clean:
	rm -f bin/$(APP_NAME)