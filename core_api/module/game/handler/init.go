package handler

import (
	"github.com/Brotiger/poker-core_api/core_api/module/game/service"
	sharedService "github.com/Brotiger/poker-core_api/core_api/shared/service"
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
