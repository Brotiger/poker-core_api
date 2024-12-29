SHELL := /bin/bash

build:
	go build -o core_api ./cmd/main.go
.PHONY: build

up:
	./core_api
.PHONY: up

run:
	go run ./cmd
.PHONY: run

dev:
	go install github.com/cespare/reflex@latest
	go install github.com/swaggo/swag/cmd/swag@latest
.PHONY: dev

docs:
	swag init -d ./cmd,./internal/module -o ./docs/swagger --parseDependency --parseInternal
.PHONY: docs