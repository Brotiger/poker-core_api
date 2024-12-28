package handler

import (
	"github.com/Brotiger/per-painted_poker-backend/app/module/game/service"
	sharedService "github.com/Brotiger/per-painted_poker-backend/app/shared/service"
)

type GameHandler struct {
	SharedTokenService *sharedService.TokenService
	GameService        *service.GameService
}

func NewGame() *GameHandler {
	return &GameHandler{
		SharedTokenService: sharedService.NewTokenService(),
		GameService:        service.NewGameService(),
	}
}
