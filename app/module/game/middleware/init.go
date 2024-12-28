package middleware

import "github.com/Brotiger/per-painted_poker-backend/app/module/game/service"

type GameMiddleware struct {
	gameService *service.GameService
}

func NewGameMiddleware() *GameMiddleware {
	return &GameMiddleware{
		gameService: service.NewGameService(),
	}
}
