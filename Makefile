.PHONY: up down build run migrate swagger test lint

## Start all containers
up:
	docker compose up -d

## Stop all containers
down:
	docker compose down

## Build the Docker image
build:
	docker compose build

## Run the API locally (no Docker)
run:
	go run ./cmd/api

## Run the queue worker locally
worker:
	go run ./cmd/worker

## Run the scheduler locally
scheduler:
	go run ./cmd/scheduler

## Run database migrations
migrate:
	go run ./cmd/migrate

## Generate Swagger docs  (requires: go install github.com/swaggo/swag/cmd/swag@latest)
swagger:
	swag init -g cmd/api/main.go -o docs

## Run all tests
test:
	go test ./... -v -race

## Run tests with coverage
test-cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

## Lint (requires: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
lint:
	golangci-lint run ./...

## Tidy dependencies
tidy:
	go mod tidy
