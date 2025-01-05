package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/connection"
	"github.com/Brotiger/poker-core_api/core_api/module/game/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerRepository struct{}

func NewPlayerRepository() *PlayerRepository {
	return &PlayerRepository{}
}

func (pr *PlayerRepository) CreatePlayer(ctx context.Context, modelPlayer model.Player) (*primitive.ObjectID, error) {
	result, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Player).InsertOne(
		ctx,
		modelPlayer,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert one, error: %w", err)
	}

	insertId := result.InsertedID.(primitive.ObjectID)
	return &insertId, nil
}

func (pr *PlayerRepository) CountPlayersByGameId(ctx context.Context, gameId primitive.ObjectID) (int64, error) {
	count, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Player).CountDocuments(
		ctx,
		bson.M{
			"gameId": gameId,
		},
		options.Count().SetHint(bson.D{
			{Key: "gameId", Value: 1},
		}),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to count documents, error: %w", err)
	}

	return count, nil
}
