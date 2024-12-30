package controller

import (
	"context"
	"errors"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	cError "github.com/Brotiger/poker-core_api/core_api/module/auth/error"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/request"
	sharedResponse "github.com/Brotiger/poker-core_api/core_api/shared/response"
	"github.com/Brotiger/poker-core_api/core_api/validator"
	"github.com/gofiber/fiber/v2"
)

// @Summary Востановление пароля
// @Tags Auth
// @Router /auth/restore [post]
// @Produce json
// @Param request body request.Restore false "Body params"
// @Success 200 {object} response.Restore "Успешный ответ."
// @Failure 400 {object} sharedResponse.Error400 "Не валидный запрос."
// @Failure 500 "Ошибка сервера."
func (ah *AuthController) Restore(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	var requestRestore request.Restore
	if err := c.BodyParser(&requestRestore); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requestRestore); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	if err := ah.AuthService.Restore(ctx, requestRestore.Email); err != nil {
		if errors.Is(err, cError.ErrUserNotFound) {
			return nil
		}

		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return nil
}
