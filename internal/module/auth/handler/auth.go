package handler

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/request"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/response"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/service"
	"github.com/Brotiger/per-painted_poker-backend/internal/validator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct {
	Service *service.Auth
}

func NewAuth() *Auth {
	return &Auth{
		Service: service.NewAuth(),
	}
}

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

	modelUser, err := a.Service.Login(ctx, requetLogin)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if modelUser == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Не верное имя пользователя или пароль.")
	}

	res, err := a.Service.GenerateTokens(ctx, modelUser.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(res)
}

func (a *Auth) Refresh(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Microsecond)
	defer cancelCtx()

	var requetRefresh request.Refresh
	if err := c.BodyParser(&requetRefresh); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error400{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetRefresh); err != nil {
		fieldErrors := validator.ValidateErr(err)

		return c.Status(fiber.StatusBadRequest).JSON(response.Error400{
			Message: "Ошибка валидации.",
			Errors:  fieldErrors,
		})
	}

	tokenClaims, err := a.Service.VerifyToken(requetRefresh.RefreshToken)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Неверный или просроченный токен обновления.")
	}

	userId, err := primitive.ObjectIDFromHex(tokenClaims.UserId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	res, err := a.Service.GenerateTokens(ctx, userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(res)
}
