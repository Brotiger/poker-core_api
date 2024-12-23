package handler

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/request"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/game/response"
	sharedResponse "github.com/Brotiger/per-painted_poker-backend/internal/shared/response"
	"github.com/Brotiger/per-painted_poker-backend/internal/validator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Создание игры
// @Tags Game
// @Router /game [post]
// @Produce json
// @Param request body request.Create false "Body params"
// @Success 200 {object} response.Create "Успешный ответ."
// @Failure 400 {object} sharedResponse.Error400 "Не валидный запрос."
// @Failure 401 {object} sharedResponse.Error401 "Невалидный токен."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (a *Game) Create(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Microsecond)
	defer cancelCtx()

	var requetCreate request.Create
	if err := c.BodyParser(&requetCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	userId := c.Locals("userId").(primitive.ObjectID)
	modelGame, err := a.GameService.CreateGame(ctx, userId, requetCreate)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(response.Create{
		Id:         modelGame.Id,
		Status:     modelGame.Status,
		Name:       modelGame.Name,
		Password:   modelGame.Password,
		OwnerId:    modelGame.OwnerId,
		Users:      modelGame.Users,
		MaxPlayers: modelGame.MaxPlayers,
	})
}
