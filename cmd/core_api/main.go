package main

import (
	"context"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/connection"
	"github.com/Brotiger/poker-core_api/core_api/router"
	"github.com/Brotiger/poker-core_api/docs/swagger"
	"github.com/Brotiger/poker-core_api/pkg/mongodb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	log "github.com/sirupsen/logrus"
)

// @title Core API
// @BasePath /api
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	ctx := context.Background()
	app := fiber.New(fiber.Config{
		DisableStartupMessage: config.Cfg.Fiber.DisableStartupMessage,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	mongodbCtx, cancelMongodbCtx := context.WithTimeout(ctx, time.Duration(config.Cfg.MongoDB.ConnectTimeoutMs)*time.Millisecond)
	defer cancelMongodbCtx()

	mongodbClient, err := mongodb.Connect(
		mongodbCtx,
		config.Cfg.MongoDB.Uri,
		config.Cfg.MongoDB.Username,
		config.Cfg.MongoDB.Password,
	)
	if err != nil {
		log.Fatalf("failed to connect to mongodb, error: %v", err)
	}

	connection.DB = mongodbClient.Database(config.Cfg.MongoDB.Database)

	options := []nats.Option{
		nats.RetryOnFailedConnect(config.Cfg.Nats.RetryOnFailedConnect),
		nats.MaxReconnects(config.Cfg.Nats.MaxReconnects),
		nats.ReconnectWait(time.Duration(config.Cfg.Nats.ReconnectWait) * time.Millisecond),
		nats.PingInterval(time.Duration(config.Cfg.Nats.PingInterval) * time.Millisecond),
		nats.MaxPingsOutstanding(config.Cfg.Nats.MaxPingOutstanding),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Warnf("nats-connect: disconnected, error: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Trace("nats-connect: reconnected")
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			if nc.Status() == nats.CLOSED && nc.LastError() != nil {
				log.Fatalf("nats-connect: connection closed after max reconnects: %v", nc.LastError())
			} else if nc.LastError() != nil {
				log.Errorf("nats-connect: connection closed, error: %v", nc.LastError())
			} else {
				log.Error("nats-connect: connection closed")
			}
		}),
	}

	if config.Cfg.Nats.ClientCert != "" && config.Cfg.Nats.ClientKey != "" {
		options = append(options, nats.ClientCert(config.Cfg.Nats.ClientCert, config.Cfg.Nats.ClientKey))
	}

	if config.Cfg.Nats.RootCA != "" {
		options = append(options, nats.RootCAs(config.Cfg.Nats.RootCA))
	}

	natsConn, err := nats.Connect(config.Cfg.Nats.Addr, options...)
	if err != nil {
		log.Fatalf("failed to connect to nats, error: %v", err)
	}

	connection.JS, err = jetstream.New(natsConn)
	if err != nil {
		log.Fatalf("failed to connect to jet stream, error: %v", err)
	}

	router.SetupRouter(app)

	log.Info("application started")
	log.Infof("local API docs: http://%s", swagger.SwaggerInfo.Host+"/swagger/index.html")

	if err := app.Listen(config.Cfg.Fiber.Listen); err != nil {
		log.Errorf("failed to listen, error: %v", err)
	}
}
