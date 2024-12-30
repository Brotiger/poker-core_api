SHELL := /bin/bash

build:
	go build -o core_api ./cmd/core_api/main.go
.PHONY: build

up:
	./core_api
.PHONY: up

run:
	go run ./cmd/core_api
.PHONY: run

dev:
	go install github.com/cespare/reflex@latest
	go install github.com/swaggo/swag/cmd/swag@latest
.PHONY: dev

docs:
	swag init -d ./cmd/core_api,./core_api/module -o ./docs/swagger --parseDependency --parseInternal
.PHONY: docs

nats-streams-add:
	nats -s "${CORE_API_NATS_ADDR}" str add "${CORE_API_NATS_STREAMS_MAILER}" --subjects "mailer" --ack --retention=work --max-age=15m --defaults;
.PHONY: nats-streams-add