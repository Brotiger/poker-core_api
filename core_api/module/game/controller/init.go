package controller

import (
	"github.com/Brotiger/poker-core_api/core_api/module/game/service"
)

type GameController struct {
	GameService  *service.GameService
	TokenService *service.TokenService
}

func NewGameController() *GameController {
	return &GameController{
		GameService:  service.NewGameService(),
		TokenService: service.NewTokenService(),
	}
}
