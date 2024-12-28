package handler

import (
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/game/service"
	sharedService "github.com/Brotiger/per-painted_poker-backend/core_api/shared/service"
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
