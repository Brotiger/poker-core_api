package router

import (
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/handler"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/middleware"
	sharedMiddleware "github.com/Brotiger/per-painted_poker-backend/internal/shared/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRouter(router fiber.Router) fiber.Router {
	gameHandler := handler.NewGame()

	authMiddleware := sharedMiddleware.NewShared()
	gameMiddleware := middleware.NewGame()

	router.Use(authMiddleware.Token, gameMiddleware.AlreadyHasGame)

	router = router.Group("/game")
	router.Get("/", gameHandler.List)
	router.Post("/start", gameHandler.Start)

	return router
}
