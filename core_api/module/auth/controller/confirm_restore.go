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
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Подтверждение кода востановления
// @Tags Auth
// @Router /auth/confirm_restore [post]
// @Produce json
// @Param request body request.ConfirmedRestore false "Body params"
// @Success 200 "Успешный ответ."
// @Failure 400 {object} sharedResponse.Error400 "Не валидный запрос."
// @Failure 500 "Ошибка сервера."
func (ah *AuthController) ConfirmRestore(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	var requestConfirmedEmail request.ConfirmedRestore
	if err := c.BodyParser(&requestConfirmedEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requestConfirmedEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	if err := ah.AuthService.ConfirmRestore(
		ctx,
		service.RequestConfirmedRestoreDTO{
			Email:    requestConfirmedEmail.Email,
			Code:     requestConfirmedEmail.Code,
			Password: requestConfirmedEmail.Password,
		},
	); err != nil {
		if errors.Is(err, cError.ErrCompareCode) || errors.Is(err, cError.ErrCodeNotFound) || errors.Is(err, cError.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(sharedResponse.BadRequest{
				Message: "Невалидный код.",
				Errors: bson.M{
					"code": "Невалидный код.",
				},
			})
		}

		log.Errorf("failed to confirmed email: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(sharedResponse.OK{
		Message: "Пароль успешно изменен.",
	})
}
