package middleware

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/app/config"
	"github.com/Brotiger/per-painted_poker-backend/app/shared/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g *Game) AlreadyHasGame(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Microsecond)
	defer cancelCtx()

	userId := c.Locals("userId").(primitive.ObjectID)
	exits, err := g.ServiceGame.UserHasGame(ctx, userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if exits {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error400{
			Message: "У пользователя уже есть игра.",
		})
	}

	return c.Next()
}
