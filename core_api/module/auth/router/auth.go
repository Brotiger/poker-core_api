package router

import (
	handler "github.com/Brotiger/poker-core_api/core_api/module/auth/controller"
	sharedMiddleware "github.com/Brotiger/poker-core_api/core_api/shared/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(api fiber.Router) {
	authController := handler.NewAuthController()
	authMiddleware := sharedMiddleware.NewShared()

	auth := api.Group("/auth")
	auth.Post("/login", authController.Login)
	auth.Post("/refresh", authController.Refresh)
	auth.Post("/logout", authMiddleware.Token, authController.Logout)
	auth.Post("/register", authController.Register)
}
