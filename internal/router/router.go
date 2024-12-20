package router

import (
	swagger "github.com/Brotiger/per-painted_poker-backend/docs/swagger"
	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/handler"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRouter(app *fiber.App) {
	authHandler := handler.NewAuth()

	swagger.SwaggerInfo.Host = config.Cfg.Swagger.Host
	swagger.SwaggerInfo.Version = config.TagVersion

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api")
	auth := api.Group("/auth")

	auth.Post("/login", authHandler.Login)
}
