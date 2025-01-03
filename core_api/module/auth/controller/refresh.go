package controller

import (
	"context"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/request"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/response"
	sharedResponse "github.com/Brotiger/poker-core_api/core_api/shared/response"
	"github.com/Brotiger/poker-core_api/core_api/validator"
	pkgService "github.com/Brotiger/poker-core_api/pkg/service"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Обновление токена
// @Tags Auth
// @Router /auth/refresh [post]
// @Produce json
// @Param request body request.Refresh false "Body params"
// @Success 200 {object} response.Token "Успешный ответ."
// @Failure 400 {object} sharedResponse.BadRequest "Не валидный запрос."
// @Failure 401 {object} sharedResponse.Unauthorized "Неверный или просроченный токен обновления."
// @Failure 500 "Ошибка сервера."
func (a *AuthController) Refresh(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	var requetRefresh request.Refresh
	if err := c.BodyParser(&requetRefresh); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetRefresh); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	tokenClaims, err := a.TokenService.VerifyToken(requetRefresh.RefreshToken)
	if err != nil {
		if err == pkgService.ErrInvalidToken {
			return c.Status(fiber.StatusUnauthorized).JSON(sharedResponse.Unauthorized{
				Message: "Просроченный токен обновления.",
			})
		}

		log.Errorf("failed to verify token, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	userId, err := primitive.ObjectIDFromHex(tokenClaims.UserId)
	if err != nil {
		log.Errorf("failed to get user id, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	exist, err := a.RefreshTokenService.CheckRefreshToken(ctx, requetRefresh.RefreshToken)
	if err != nil {
		log.Errorf("failed to check refresh token, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	if !exist {
		return c.Status(fiber.StatusUnauthorized).JSON(sharedResponse.Unauthorized{
			Message: "Невалидный токен обновления.",
		})
	}

	dtoToken, err := a.RefreshTokenService.GenerateTokens(ctx, userId)
	if err != nil {
		log.Errorf("failed to generate tokens, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(response.Token{
		AccessToken:  dtoToken.AccessToken,
		RefreshToken: dtoToken.RefreshToken,
	})
}
