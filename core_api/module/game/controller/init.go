package controller

import (
	"github.com/Brotiger/poker-core_api/core_api/module/game/service"
	pkgService "github.com/Brotiger/poker-core_api/pkg/service"
)

type GameController struct {
	SharedTokenService *pkgService.TokenService
	GameService        *service.GameService
	TokenService       *service.TokenService
}

func NewGameController() *GameController {
	return &GameController{
		SharedTokenService: pkgService.NewTokenService(),
		GameService:        service.NewGameService(),
		TokenService:       service.NewTokenService(),
	}
}
