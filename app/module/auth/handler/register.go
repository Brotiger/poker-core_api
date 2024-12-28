package handler

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/app/config"
	"github.com/Brotiger/per-painted_poker-backend/app/module/auth/request"
	"github.com/Brotiger/per-painted_poker-backend/app/module/auth/service"
	sharedResponse "github.com/Brotiger/per-painted_poker-backend/app/shared/response"
	"github.com/Brotiger/per-painted_poker-backend/app/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (ah *AuthHandler) Register(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	var requetRegister request.Register
	if err := c.BodyParser(&requetRegister); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetRegister); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	if err := ah.AuthService.Register(ctx, service.RequestRegisterDTO{
		Email:    requetRegister.Email,
		Username: requetRegister.Username,
		Password: requetRegister.Password,
	}); err != nil {
		log.Errorf("failed to register, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return nil
}
