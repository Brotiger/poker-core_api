package config

import (
	"log"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	App     *App
	MongoDB *MongoDB
}

var Cfg *Config

func init() {
	appConfig := &App{}
	mongodbConfig := &MongoDB{}

	if err := env.Parse(appConfig); err != nil {
		log.Fatalf("failed to parse app config, error: %v", err)
	}

	if err := env.Parse(mongodbConfig); err != nil {
		log.Fatalf("failed to parse mongodb config, error: %v", err)
	}

	Cfg = &Config{
		App:     appConfig,
		MongoDB: mongodbConfig,
	}
}
