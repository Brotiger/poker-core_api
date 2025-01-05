package router

import (
	handler "github.com/Brotiger/poker-core_api/core_api/module/game/controller"
	sharedMiddleware "github.com/Brotiger/poker-core_api/core_api/shared/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(api fiber.Router) {
	gameController := handler.NewGameController()

	authMiddleware := sharedMiddleware.NewAuthMiddleware()

	game := api.Group("/game")
	game.Use(authMiddleware.Token)
	game.Get("", gameController.List)
	game.Post("", gameController.Create)
	game.Post("/join/:id<len(24)>", gameController.Join)
}
