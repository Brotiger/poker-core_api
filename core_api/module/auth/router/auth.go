package router

import (
	"github.com/Brotiger/poker-core_api/core_api/module/auth/handler"
	sharedMiddleware "github.com/Brotiger/poker-core_api/core_api/shared/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(api fiber.Router) {
	authHandler := handler.NewAuthHandler()
	authMiddleware := sharedMiddleware.NewShared()

	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh", authHandler.Refresh)
	auth.Post("/logout", authMiddleware.Token, authHandler.Logout)
}
