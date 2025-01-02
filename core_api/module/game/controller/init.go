package controller

import (
	"github.com/Brotiger/poker-core_api/core_api/module/game/service"
	sharedService "github.com/Brotiger/poker-core_api/core_api/shared/service"
)

type GameController struct {
	SharedTokenService *sharedService.TokenService
	GameService        *service.GameService
	TokenService       *service.TokenService
}

func NewGameController() *GameController {
	return &GameController{
		SharedTokenService: sharedService.NewTokenService(),
		GameService:        service.NewGameService(),
		TokenService:       service.NewTokenService(),
	}
}
