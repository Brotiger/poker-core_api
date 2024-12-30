package controller

import (
	"context"
	"errors"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	cError "github.com/Brotiger/poker-core_api/core_api/module/auth/error"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/request"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/service"
	sharedResponse "github.com/Brotiger/poker-core_api/core_api/shared/response"
	"github.com/Brotiger/poker-core_api/core_api/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// @Summary Подтверждение почты
// @Tags Auth
// @Router /auth/confirm_email [post]
// @Produce json
// @Param request body request.ConfirmedEmail false "Body params"
// @Success 200 "Успешный ответ."
// @Failure 400 {object} sharedResponse.Error400 "Не валидный запрос."
// @Failure 500 "Ошибка сервера."
func (ah *AuthController) ConfirmEmail(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	var requestConfirmedEmail request.ConfirmedEmail
	if err := c.BodyParser(&requestConfirmedEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requestConfirmedEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	if err := ah.AuthService.ConfirmEmail(
		ctx,
		service.RequestConfirmedEmailDTO{
			UserId: requestConfirmedEmail.UserId,
			Code:   requestConfirmedEmail.Code,
		},
	); err != nil {
		if errors.Is(err, cError.ErrCodeNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(sharedResponse.Error404{
				Message: "Код не найден.",
			})
		}
		if errors.Is(err, cError.ErrCompareCode) {
			return c.Status(fiber.StatusNotFound).JSON(sharedResponse.Error400{
				Message: "Невалидный код.",
			})
		}

		log.Errorf("failed to confirmed email: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return nil
}
