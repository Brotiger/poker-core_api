package handler

import (
	"context"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/request"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/response"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/service"
	sharedResponse "github.com/Brotiger/poker-core_api/core_api/shared/response"
	"github.com/Brotiger/poker-core_api/core_api/validator"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Регистрация
// @Tags Auth
// @Router /auth/register [post]
// @Produce json
// @Param request body request.Register false "Body params"
// @Success 200 {object} response.Register "Успешный ответ."
// @Failure 400 {object} sharedResponse.Error400 "Не валидный запрос."
// @Failure 500 "Ошибка сервера."
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

	ok, err := ah.AuthService.CheckUsername(ctx, requetRegister.Username)
	if err != nil {
		log.Errorf("failed to check username, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Ошибка валидации.",
			Errors: bson.M{
				"username": "Пользователь с таким именем уже существует.",
			},
		})
	}

	ok, err = ah.AuthService.CheckEmail(ctx, requetRegister.Email)
	if err != nil {
		log.Errorf("failed to check email, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Ошибка валидации.",
			Errors: bson.M{
				"email": "Пользователь с такой почтой уже существует.",
			},
		})
	}

	responseRegisterDTO, err := ah.AuthService.Register(ctx, service.RequestRegisterDTO{
		Email:    requetRegister.Email,
		Username: requetRegister.Username,
		Password: requetRegister.Password,
	})
	if err != nil {
		log.Errorf("failed to register, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(response.Register{
		Id: responseRegisterDTO.Id,
	})
}
