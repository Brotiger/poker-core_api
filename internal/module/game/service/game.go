package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/model"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/repository"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	GameRepository *repository.Game
}

func NewGame() *Game {
	return &Game{
		GameRepository: repository.NewGame(),
	}
}

func (g *Game) GetList(ctx context.Context, requetList request.List) ([]model.Game, int64, error) {
	modelGames, err := g.GameRepository.GetGames(ctx, requetList)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get games, error: %w", err)
	}

	total, err := g.GameRepository.GetGameCount(ctx, requetList)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get game count, error: %w", err)
	}

	return modelGames, total, nil
}

func (g *Game) Create(ctx context.Context, userId primitive.ObjectID, requestCreate request.Create) (*model.Game, error) {
	timeNow := time.Now()
	modelGame := model.Game{
		Status:     "created",
		Name:       requestCreate.Name,
		Password:   requestCreate.Password,
		MaxPlayers: requestCreate.MaxPlayers,
		Users:      []primitive.ObjectID{userId},
		UpdatedAt:  timeNow,
		CreatedAt:  timeNow,
	}

	insertId, err := g.GameRepository.CreateGame(ctx, modelGame)
	if err != nil {
		return nil, fmt.Errorf("failed to create game, error: %w", err)
	}

	modelGame.Id = insertId

	return &modelGame, nil
}

func (g *Game) UserHasGame(ctx context.Context, userId primitive.ObjectID) (bool, error) {
	count, err := g.GameRepository.CountUserGames(ctx, userId)
	if err != nil {
		return false, fmt.Errorf("failed to count user games count, error: %w", err)
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}
