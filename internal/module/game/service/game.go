package service

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/model"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/repository"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/request"
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
