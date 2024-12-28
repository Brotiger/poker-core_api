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

func (gm *GameMiddleware) GameCannotBeStarted(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	userId, err := primitive.ObjectIDFromHex(c.Locals("userId").(string))
	if err != nil {
		log.Errorf("failed to convert userId to ObjectID, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	canBeStarted, err := gm.gameService.GameCanBeStarted(ctx, userId)
	if err != nil {
		log.Errorf("failed to check if game can be started, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	if !canBeStarted {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error400{
			Message: "Игра не может быть запущена.",
		})
	}

	return c.Next()
}
