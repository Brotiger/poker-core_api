package middleware

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/shared/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g *Game) GameCannotBeStarted(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Microsecond)
	defer cancelCtx()

	userId := c.Locals("userId").(primitive.ObjectID)
	canBeStarted, err := g.ServiceGame.GameCanBeStarted(ctx, userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if !canBeStarted {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error400{
			Message: "Игра не может быть запущена.",
		})
	}

	return c.Next()
}
