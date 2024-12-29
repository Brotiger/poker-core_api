package main

import (
	"context"
	"time"

	"github.com/Brotiger/poker-core_api/pkg/mongodb"
	"github.com/Brotiger/poker-core_api/seeder/config"
	"github.com/Brotiger/poker-core_api/seeder/connection"
	"github.com/Brotiger/poker-core_api/seeder/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

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

	userService := service.NewUserService()
	if err := userService.CreateUser(ctx); err != nil {
		log.Fatalf("failed to create user, error: %v", err)
	}
}
