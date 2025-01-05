package controller

import (
	"github.com/Brotiger/poker-core_api/core_api/module/game/service"
)

type GameController struct {
	gameService         *service.GameService
	connectTokenService *service.ConnectTokenService
}

func NewGameController() *GameController {
	return &GameController{
		gameService:         service.NewGameService(),
		connectTokenService: service.NewConnectTokenService(),
	}
}
