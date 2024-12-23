package handler

import (
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/service"
	sharedService "github.com/Brotiger/per-painted_poker-backend/internal/shared/service"
)

type Game struct {
	SharedTokenService *sharedService.Token
	GameService        *service.Game
}

func NewGame() *Game {
	return &Game{
		SharedTokenService: sharedService.NewToken(),
		GameService:        service.NewGame(),
	}
}
