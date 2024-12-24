package router

import (
	"github.com/Brotiger/per-painted_poker-backend/app/config"
	authRouter "github.com/Brotiger/per-painted_poker-backend/app/module/auth/router"
	swagger "github.com/Brotiger/per-painted_poker-backend/docs/swagger"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRouter(app *fiber.App) *fiber.Router {
	swagger.SwaggerInfo.Host = config.Cfg.Fiber.Swagger.Host
	swagger.SwaggerInfo.Version = config.TagVersion

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	router := app.Group("/api")
	router = authRouter.SetupAuthRouter(router)

	return &router
}
