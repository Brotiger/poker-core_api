package router

import (
	"github.com/Brotiger/per-painted_poker-backend/app/module/auth/handler"
	sharedMiddleware "github.com/Brotiger/per-painted_poker-backend/app/shared/middleware"
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
