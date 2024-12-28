SHELL := /bin/bash

build:
	go build -o app ./cmd/core_api/main.go
.PHONY: build

up:
	./core_api
.PHONY: up

run:
	go run ./cmd/core_api
.PHONY: run

seed:
	go run ./cmd/seeder
.PHONY: seed

dev:
	go install github.com/cespare/reflex@latest
	go install github.com/swaggo/swag/cmd/swag@latest
.PHONY: dev

docs:
	swag init -d ./cmd/core_api,./core_api/module -o ./docs/swagger --parseDependency --parseInternal
.PHONY: docs