package handler

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/request"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/response"
	"github.com/Brotiger/per-painted_poker-backend/internal/validator"
	"github.com/gofiber/fiber/v2"
)

// @Summary Авторизация
// @Tags Auth
// @Router /auth/login [post]
// @Produce json
// @Failure 200 {object} response.Login "Успешный ответ."
// @Failure 400 {object} response.Error400 "Не валидный запрос."
// @Failure 401 "Не верное имя пользователя или пароль."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (a *Auth) Login(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Microsecond)
	defer cancelCtx()

	var requetLogin request.Login
	if err := c.BodyParser(&requetLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error400{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetLogin); err != nil {
		fieldErrors := validator.ValidateErr(err)

		return c.Status(fiber.StatusBadRequest).JSON(response.Error400{
			Message: "Ошибка валидации.",
			Errors:  fieldErrors,
		})
	}

	modelUser, err := a.AuthService.Login(ctx, requetLogin)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if modelUser == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Не верное имя пользователя или пароль.")
	}

	res, err := a.RefreshTokenService.GenerateTokens(ctx, modelUser.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(res)
}
