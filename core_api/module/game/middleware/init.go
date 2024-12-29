package middleware

import "github.com/Brotiger/poker-core_api/core_api/module/game/service"

type GameMiddleware struct {
	gameService *service.GameService
}

func NewGameMiddleware() *GameMiddleware {
	return &GameMiddleware{
		gameService: service.NewGameService(),
	}
}
