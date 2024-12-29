package router

import (
	"github.com/Brotiger/poker-core_api/core_api/config"
	authRouter "github.com/Brotiger/poker-core_api/core_api/module/auth/router"
	gameRouter "github.com/Brotiger/poker-core_api/core_api/module/game/router"
	swagger "github.com/Brotiger/poker-core_api/docs/swagger"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRouter(app *fiber.App) {
	swagger.SwaggerInfo.Host = config.Cfg.Fiber.Swagger.Host
	swagger.SwaggerInfo.Version = config.TagVersion

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api")
	authRouter.SetupRouter(api)
	gameRouter.SetupRouter(api)
}
