package controller

import (
	"context"
	"errors"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	cError "github.com/Brotiger/poker-core_api/core_api/module/auth/error"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/request"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/response"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/service"
	sharedResponse "github.com/Brotiger/poker-core_api/core_api/shared/response"
	"github.com/Brotiger/poker-core_api/core_api/validator"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// @Summary Авторизация
// @Tags Auth
// @Router /auth/login [post]
// @Produce json
// @Param request body request.Login false "Body params"
// @Success 200 {object} response.Token "Успешный ответ."
// @Failure 400 {object} sharedResponse.BadRequest "Не валидный запрос."
// @Failure 403 {object} sharedResponse.Forbidden "Не верное имя пользователя или пароль."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (a *AuthController) Login(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	var requestLogin request.Login
	if err := c.BodyParser(&requestLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requestLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	modelUser, err := a.AuthService.GetUser(ctx, service.RequestGetUserDTO{
		Email:    requestLogin.Email,
		Password: requestLogin.Password,
	})
	if err != nil {
		if errors.Is(err, cError.ErrUserNotFound) || errors.Is(err, cError.ErrCompareHashAndPassword) {
			return c.Status(fiber.StatusForbidden).JSON(sharedResponse.Forbidden{
				Message: "Не верное имя пользователя или пароль.",
			})
		}

		log.Errorf("failed to get user, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	if !modelUser.EmailConfirmed {
		return c.Status(fiber.StatusUnauthorized).JSON(sharedResponse.Unauthorized{
			Message: "Пользователь не подтвержден.",
		})
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
