package handler

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/app/config"
	"github.com/Brotiger/per-painted_poker-backend/app/module/game/request"
	"github.com/Brotiger/per-painted_poker-backend/app/module/game/response"
	sharedResponse "github.com/Brotiger/per-painted_poker-backend/app/shared/response"
	"github.com/Brotiger/per-painted_poker-backend/app/validator"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
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
		log.Errorf("failed to create game, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
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
