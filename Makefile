SHELL := /bin/bash

build:
	go build -o app ./cmd/app/main.go
.PHONY: build

up:
	./app
.PHONY: up

run:
	go run ./cmd/app
.PHONY: run

seed:
	go run ./cmd/seeder
.PHONY: seed

dev:
	go install github.com/cespare/reflex@latest
	go install github.com/swaggo/swag/cmd/swag@latest
.PHONY: dev

docs:
	swag init -d ./cmd,./app,./app/module -o ./docs/swagger --parseDependency --parseInternal
.PHONY: docs