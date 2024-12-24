package router

import (
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/handler"
	sharedMiddleware "github.com/Brotiger/per-painted_poker-backend/internal/shared/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRouter(router fiber.Router) fiber.Router {
	authHandler := handler.NewAuth()

	authMiddleware := sharedMiddleware.NewShared()

	router = router.Group("/auth")
	router.Post("/login", authHandler.Login)
	router.Post("/refresh", authHandler.Refresh)

	router.Use(authMiddleware.Token)
	router.Post("/logout", authHandler.Logout)

	return router
}
