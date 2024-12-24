package handler

import (
	"github.com/Brotiger/per-painted_poker-backend/app/module/game/service"
	sharedService "github.com/Brotiger/per-painted_poker-backend/app/shared/service"
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
