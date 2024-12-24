package handler

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/app/config"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Выход
// @Tags Auth
// @Router /auth/logout [post]
// @Produce json
// @Success 200 "Успешный ответ."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (a *Auth) Logout(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	userId, err := primitive.ObjectIDFromHex(*(c.Locals("userId")).(*string))
	if err != nil {
		log.Errorf("failed to get user id, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	if err := a.RefreshTokenService.Logout(ctx, userId); err != nil {
		log.Errorf("failed to logout user, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
