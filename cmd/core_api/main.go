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

	router.SetupRouter(app)

	log.Info("application started")
	log.Infof("local API docs: http://%s", swagger.SwaggerInfo.Host+"/swagger/index.html")

	if err := app.Listen(config.Cfg.Fiber.Listen); err != nil {
		log.Errorf("failed to listen, error: %v", err)
	}
}
