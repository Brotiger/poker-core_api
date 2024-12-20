package router

import (
	"github.com/Brotiger/per-painted_poker-backend/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	auth := handler.NewAuth()
	app.Post("/login", auth.Login)
}
