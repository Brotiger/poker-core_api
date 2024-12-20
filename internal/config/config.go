package config

import (
	"log"

	"github.com/caarlos0/env"
)

const TagVersion = "0.0.1"

type Config struct {
	App     *App
	Fiber   *Fiber
	MongoDB *MongoDB
	Table   *Table
	Swagger *Swagger
}

var Cfg Config

func init() {
	appConfig := &App{}
	fiberConfig := &Fiber{}
	mongodbConfig := &MongoDB{}
	tableConfig := &Table{}
	swaggerConfig := &Swagger{}

	if err := env.Parse(appConfig); err != nil {
		log.Fatalf("failed to parse app config, error: %v", err)
	}

	if err := env.Parse(fiberConfig); err != nil {
		log.Fatalf("failed to parse fiber config, error: %v", err)
	}

	if err := env.Parse(mongodbConfig); err != nil {
		log.Fatalf("failed to parse mongodb config, error: %v", err)
	}

	if err := env.Parse(tableConfig); err != nil {
		log.Fatalf("failed to parse table config, error: %v", err)
	}

	if err := env.Parse(swaggerConfig); err != nil {
		log.Fatalf("failed to parse swagger config, error: %v", err)
	}

	Cfg = Config{
		App:     appConfig,
		Fiber:   fiberConfig,
		MongoDB: mongodbConfig,
		Table:   tableConfig,
		Swagger: swaggerConfig,
	}
}
