package router

import (
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRouter(router fiber.Router) fiber.Router {
	authHandler := handler.NewAuth()
	authRouter := router.Group("/auth")
	authRouter.Post("/login", authHandler.Login)
	authRouter.Post("/refresh", authHandler.Refresh)

	return authRouter
}
