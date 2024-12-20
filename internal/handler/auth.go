package handler

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/request"
	"github.com/Brotiger/per-painted_poker-backend/internal/response"
	"github.com/Brotiger/per-painted_poker-backend/internal/service"
	"github.com/Brotiger/per-painted_poker-backend/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	Service *service.Auth
}

func NewAuth() *Auth {
	return &Auth{
		Service: service.NewAuth(),
	}
}

func (a *Auth) Login(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Microsecond)
	defer cancelCtx()

	var requetLogin request.Login

	if err := c.BodyParser(&requetLogin); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validator.Struct(requetLogin); err != nil {
		fieldErrors := validator.ValidateErr(err)
		res := response.Error400{
			Message: "Ошибка валидации",
			Errors:  fieldErrors,
		}

		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	token, err := a.Service.Login(ctx, requetLogin)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if token == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid username or password")
	}

	accessToken, err := token.SignedString([]byte(config.Cfg.App.Jwt.Secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(response.Login{
		AccessToken: accessToken,
	})
}
