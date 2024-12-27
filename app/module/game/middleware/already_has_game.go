package middleware

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/app/config"
	"github.com/Brotiger/per-painted_poker-backend/app/shared/response"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g *Game) AlreadyHasGame(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	userId, err := primitive.ObjectIDFromHex(c.Locals("userId").(string))
	if err != nil {
		log.Errorf("failed to convert userId to ObjectID, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	exits, err := g.ServiceGame.UserHasGame(ctx, userId)
	if err != nil {
		log.Errorf("failed to check user has game, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	if exits {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error400{
			Message: "У пользователя уже есть игра.",
		})
	}

	return c.Next()
}
