SHELL := /bin/bash

build:
	go build -o main ./cmd/main.go
.PHONY: build

up:
	./main
.PHONY: up

run:
	go run ./cmd
.PHONY: run

dev:
	go install github.com/cespare/reflex@latest
	go install github.com/swaggo/swag/cmd/swag@latest
.PHONY: dev

docs:
	swag init -d ./cmd,./internal,./internal/module -o ./docs/swagger --parseDependency --parseInternal
.PHONY: docs