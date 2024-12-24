package middleware

import "github.com/Brotiger/per-painted_poker-backend/internal/module/game/service"

type Game struct {
	ServiceGame *service.Game
}

func NewGame() *Game {
	return &Game{
		ServiceGame: service.NewGame(),
	}
}
