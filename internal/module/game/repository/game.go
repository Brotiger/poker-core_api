package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/connection"
	cError "github.com/Brotiger/per-painted_poker-backend/internal/module/game/error"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/model"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/request"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) GetGames(ctx context.Context, request request.List) ([]model.Game, error) {
	hint := bson.D{
		{Key: "createdAt", Value: 1},
	}

	filter := bson.M{}

	if request.Name != "" {
		filter["name"] = request.Name
		hint = append(hint, bson.E{Key: "name", Value: 1})
	}

	cur, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Game).Find(
		ctx,
		filter,
		options.Find().SetSkip(request.From).SetLimit(request.Size).SetSort(bson.M{"createdAt": 1}).SetHint(hint),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find, error: %w", err)
	}

	defer func() {
		if err := cur.Close(ctx); err != nil {
			log.Errorf("failed to close cursor, error: %v", err)
		}
	}()

	modelGames := []model.Game{}
	for cur.Next(ctx) {
		var modelGame model.Game
		if err := cur.Decode(&modelGame); err != nil {
			return nil, fmt.Errorf("failed to decode game to model, error: %w", err)
		}

		modelGames = append(modelGames, modelGame)
	}

	if err = cur.Err(); err != nil {
		return nil, fmt.Errorf("failed to process mongodb, cursor error: %w", err)
	}

	return modelGames, nil
}

func (g *Game) GetGameCount(ctx context.Context, request request.List) (int64, error) {
	hint := bson.D{}
	filter := bson.M{}

	if request.Name != "" {
		filter["name"] = request.Name
		hint = append(hint, bson.E{Key: "name", Value: 1})
	}

	count, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Game).CountDocuments(
		ctx,
		filter,
		options.Count().SetHint(hint),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to count documents, error: %w", err)
	}

	return count, nil
}

func (g *Game) CreateGame(ctx context.Context, modelGame model.Game) (primitive.ObjectID, error) {
	var inserId primitive.ObjectID
	result, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Game).InsertOne(
		ctx,
		modelGame,
	)
	if err != nil {
		return inserId, fmt.Errorf("failed to insert one, error: %w", err)
	}

	inserId = result.InsertedID.(primitive.ObjectID)
	return inserId, nil
}

func (g *Game) CountUserGames(ctx context.Context, ownerId primitive.ObjectID) (int64, error) {
	count, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Game).CountDocuments(
		ctx,
		bson.M{"ownerId": ownerId},
		options.Count().SetHint(bson.D{
			{Key: "ownerId", Value: 1},
		}),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to count documents, error: %w", err)
	}

	return count, nil
}

func (g *Game) UpdateGameStatus(ctx context.Context, userId primitive.ObjectID, status string) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Game).UpdateOne(
		ctx,
		bson.M{"ownerId": userId},
		bson.M{
			"$set": bson.M{
				"status": status,
			},
		},
		options.Update().SetHint(bson.D{
			{Key: "ownerId", Value: 1},
		}),
	); err != nil {
		return fmt.Errorf("failed to update one, error: %w", err)
	}

	return nil
}

func (g *Game) GetGameByOwnerId(ctx context.Context, ownerId primitive.ObjectID) (*model.Game, error) {
	var modelGame model.Game
	if err := connection.DB.Collection(config.Cfg.MongoDB.Table.Game).FindOne(
		ctx,
		bson.M{
			"ownerId": ownerId,
		},
		options.FindOne().SetHint(bson.D{
			{Key: "ownerId", Value: 1},
		}),
	).Decode(&modelGame); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, cError.ErrGameNotFound
		}

		return nil, fmt.Errorf("failed to find one, error: %w", err)
	}

	return &modelGame, nil
}
