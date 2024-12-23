package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/connection"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/model"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/request"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) GetGames(ctx context.Context, request request.List) ([]model.Game, error) {
	filter := bson.M{}
	if request.Name != "" {
		filter["name"] = request.Name
	}

	cur, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Game).Find(
		ctx,
		filter,
		options.Find().SetSkip(request.From).SetLimit(request.Size),
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
	filter := bson.M{}
	if request.Name != "" {
		filter["name"] = request.Name
	}

	count, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Game).CountDocuments(
		ctx,
		filter,
		options.Count(),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to count documents, error: %w", err)
	}

	return count, nil
}