package router

import (
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/handler"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/middleware"
	sharedMiddleware "github.com/Brotiger/per-painted_poker-backend/internal/shared/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRouter(router fiber.Router) fiber.Router {
	gameHandler := handler.NewGame()
	router = router.Group("/game")

	router.Use(sharedMiddleware.Token, middleware.AlreadyHasGame)
	router.Get("/", gameHandler.List)

	return router
}
