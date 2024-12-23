package handler

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/request"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/response"
	sharedResponse "github.com/Brotiger/per-painted_poker-backend/internal/shared/response"
	sharedService "github.com/Brotiger/per-painted_poker-backend/internal/shared/service"
	"github.com/Brotiger/per-painted_poker-backend/internal/validator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Обновление токена
// @Tags Auth
// @Router /auth/refresh [post]
// @Produce json
// @Param request body request.Refresh false "Body params"
// @Success 200 {object} response.Token "Успешный ответ."
// @Failure 400 {object} sharedResponse.Error400 "Не валидный запрос."
// @Failure 401 {object} sharedResponse.Error401 "Неверный или просроченный токен обновления."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (a *Auth) Refresh(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Microsecond)
	defer cancelCtx()

	var requetRefresh request.Refresh
	if err := c.BodyParser(&requetRefresh); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetRefresh); err != nil {
		fieldErrors := validator.ValidateErr(err)

		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Ошибка валидации.",
			Errors:  fieldErrors,
		})
	}

	tokenClaims, err := a.SharedTokenService.VerifyToken(requetRefresh.RefreshToken)
	if err != nil {
		if err == sharedService.ErrInvalidToken {
			return c.Status(fiber.StatusUnauthorized).JSON(sharedResponse.Error401{
				Message: "Просроченный токен обновления.",
			})
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	userId, err := primitive.ObjectIDFromHex(tokenClaims.UserId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	exist, err := a.RefreshTokenService.CheckRefreshToken(ctx, userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if !exist {
		return c.Status(fiber.StatusUnauthorized).JSON(sharedResponse.Error401{
			Message: "Невалидный токен обновления.",
		})
	}

	dtoToken, err := a.RefreshTokenService.GenerateTokens(ctx, userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(response.Token{
		AccessToken:  dtoToken.AccessToken,
		RefreshToken: dtoToken.RefreshToken,
	})
}
