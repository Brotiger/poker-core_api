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

// @Summary Создание игры
// @Tags Game
// @Router /game [post]
// @Produce json
// @Param request body request.Create false "Body params"
// @Success 200 {object} response.Create "Успешный ответ."
// @Failure 400 {object} sharedResponse.BadRequest "Не валидный запрос."
// @Failure 401 {object} sharedResponse.Unauthorized "Невалидный токен."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (gh *GameController) Create(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	var requetCreate request.Create
	if err := c.BodyParser(&requetCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetCreate); err != nil {
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

	responsCreateGameDTO, err := gh.gameService.CreateGame(ctx, service.RequestCreateGameDTO{
		UserId:     userId,
		Name:       requetCreate.Name,
		Password:   requetCreate.Password,
		MaxPlayers: requetCreate.MaxPlayers,
	})
	if err != nil {
		log.Errorf("failed to create game, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	connectToken, err := gh.connectTokenService.GenerateConnectToken(ctx, service.RequestGenerateConnectToken{
		GameId: responsCreateGameDTO.Id,
		UserId: userId,
	})
	if err != nil {
		log.Errorf("failed to generate connect token, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(response.Create{
		Game: response.CreateGame{
			Id:           responsCreateGameDTO.Id,
			Status:       responsCreateGameDTO.Status,
			Name:         responsCreateGameDTO.Name,
			WithPassword: responsCreateGameDTO.Password != nil,
			MaxPlayers:   responsCreateGameDTO.MaxPlayers,
		},
		ConnectToken: connectToken,
	})
}
