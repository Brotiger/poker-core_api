package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env/v7"
)

const TagVersion = "0.0.1"

type Config struct {
	App     *App
	Fiber   *Fiber
	MongoDB *MongoDB
	Nats    *Nats
}

var Cfg *Config

func init() {
	appConfig := &App{}
	fiberConfig := &Fiber{}
	mongodbConfig := &MongoDB{}
	natsConfig := &Nats{}

	if err := env.Parse(appConfig); err != nil {
		log.Fatalf("failed to parse app config, error: %v", err)
	}

	if err := env.Parse(fiberConfig); err != nil {
		log.Fatalf("failed to parse fiber config, error: %v", err)
	}

	if err := env.Parse(mongodbConfig); err != nil {
		log.Fatalf("failed to parse mongodb config, error: %v", err)
	}

	if err := env.Parse(natsConfig); err != nil {
		log.Fatalf("failed to parse nats config, error: %v", err)
	}

	Cfg = &Config{
		App:     appConfig,
		Fiber:   fiberConfig,
		MongoDB: mongodbConfig,
		Nats:    natsConfig,
	}
}
