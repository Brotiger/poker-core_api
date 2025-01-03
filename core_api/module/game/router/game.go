package router

import (
	handler "github.com/Brotiger/poker-core_api/core_api/module/game/controller"
	"github.com/Brotiger/poker-core_api/core_api/module/game/middleware"
	sharedMiddleware "github.com/Brotiger/poker-core_api/core_api/shared/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(api fiber.Router) {
	gameController := handler.NewGameController()

	authMiddleware := sharedMiddleware.NewAuthMiddleware()
	gameMiddleware := middleware.NewGameMiddleware()

	game := api.Group("/game")
	game.Use(authMiddleware.Token)
	game.Get("", gameController.List)
	game.Post("", gameMiddleware.AlreadyHasGame, gameController.Create)
	game.Post("/start", gameController.Start)
	game.Post("/join", gameController.Join)
}
