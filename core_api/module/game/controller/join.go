package controller

import (
	"context"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/module/game/request"
	"github.com/Brotiger/poker-core_api/core_api/module/game/response"
	"github.com/Brotiger/poker-core_api/core_api/module/game/service"
	sharedResponse "github.com/Brotiger/poker-core_api/core_api/shared/response"
	"github.com/Brotiger/poker-core_api/core_api/validator"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Подключение к игре
// @Tags Game
// @Router /game/join [post]
// @Produce json
// @Param request body request.Join false "Body params"
// @Success 200 {object} response.Join "Успешный ответ."
// @Failure 400 {object} sharedResponse.BadRequest "Не валидный запрос."
// @Failure 401 {object} sharedResponse.Unauthorized "Невалидный токен."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (gh *GameController) Join(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	var requetJoin request.Join
	if err := c.BodyParser(&requetJoin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetJoin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	userId, err := primitive.ObjectIDFromHex(c.Locals("userId").(string))
	if err != nil {
		log.Errorf("failed to convert userId to ObjectID, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	modelGame, err := gh.GameService.JoinGame(ctx, service.RequestJoinGameDTO{
		UserId:   userId,
		GameId:   requetJoin.GameId,
		Password: requetJoin.Password,
	})
	if err != nil {
		log.Errorf("failed to create game, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	connectToken, err := gh.TokenService.GenerateConnectToken(ctx, modelGame.Id)
	if err != nil {
		log.Errorf("failed to generate connect token, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(response.Join{
		Game: response.CreateGame{
			Id:           modelGame.Id,
			Status:       modelGame.Status,
			Name:         modelGame.Name,
			WithPassword: modelGame.Password != nil,
			MaxPlayers:   modelGame.MaxPlayers,
			Users:        modelGame.Users,
		},
		ConnectToken: connectToken,
	})
}
