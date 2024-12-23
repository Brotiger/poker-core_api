package router

import (
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/handler"
	"github.com/Brotiger/per-painted_poker-backend/internal/shared/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRouter(router fiber.Router) fiber.Router {
	gameHandler := handler.NewGame()
	router = router.Group("/game")

	router.Use(middleware.Token)
	router.Get("/list", gameHandler.List)

	return router
}
