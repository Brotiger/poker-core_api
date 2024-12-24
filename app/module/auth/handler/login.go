package handler

import (
	"context"
	"errors"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/app/config"
	cError "github.com/Brotiger/per-painted_poker-backend/app/module/auth/error"
	"github.com/Brotiger/per-painted_poker-backend/app/module/auth/request"
	"github.com/Brotiger/per-painted_poker-backend/app/module/auth/response"
	sharedResponse "github.com/Brotiger/per-painted_poker-backend/app/shared/response"
	"github.com/Brotiger/per-painted_poker-backend/app/validator"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// @Summary Авторизация
// @Tags Auth
// @Router /auth/login [post]
// @Produce json
// @Param request body request.Login false "Body params"
// @Success 200 {object} response.Token "Успешный ответ."
// @Failure 400 {object} sharedResponse.Error400 "Не валидный запрос."
// @Failure 401 {object} sharedResponse.Error401 "Не верное имя пользователя или пароль."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (a *Auth) Login(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Microsecond)
	defer cancelCtx()

	var requetLogin request.Login
	if err := c.BodyParser(&requetLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	modelUser, err := a.AuthService.GetUser(ctx, requetLogin)
	if err != nil {
		if errors.Is(err, cError.ErrUserNotFound) || errors.Is(err, cError.ErrCompareHashAndPassword) {
			return c.Status(fiber.StatusUnauthorized).JSON(sharedResponse.Error401{
				Message: "Не верное имя пользователя или пароль.",
			})
		}

		log.Errorf("failed to get user, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	dtoToken, err := a.RefreshTokenService.GenerateTokens(ctx, modelUser.Id)
	if err != nil {
		log.Errorf("failed to generate tokens, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(response.Token{
		AccessToken:  dtoToken.AccessToken,
		RefreshToken: dtoToken.RefreshToken,
	})
}
