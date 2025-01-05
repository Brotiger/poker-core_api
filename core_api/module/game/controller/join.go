package controller

import (
	"context"
	"errors"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	cError "github.com/Brotiger/poker-core_api/core_api/module/game/error"
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
// @Router /game/join/{id} [post]
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

	gameId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		log.Errorf("failed to convert gameId to ObjectID, error: %v", err)
		return fiber.NewError(fiber.StatusBadRequest)
	}

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

	allow, err := gh.gameService.CheckGameAllowToJoin(ctx, gameId)
	if err != nil {
		log.Errorf("failed to check game allow to join, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	if !allow {
		return c.Status(fiber.StatusForbidden).JSON(sharedResponse.Forbidden{
			Message: "Игра не доступна для подключения.",
		})
	}

	userId, err := primitive.ObjectIDFromHex(c.Locals("userId").(string))
	if err != nil {
		log.Errorf("failed to convert userId to ObjectID, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	responsJoinGameDTOUsers, err := gh.gameService.JoinGame(ctx, service.RequestJoinGameDTO{
		UserId:   userId,
		GameId:   gameId,
		Password: requetJoin.Password,
	})
	if err != nil {
		if errors.Is(err, cError.ErrComparePassword) {
			return c.Status(fiber.StatusForbidden).JSON(sharedResponse.Forbidden{
				Message: "Неверный пароль.",
			})
		}

		log.Errorf("failed to create game, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	connectToken, err := gh.connectTokenService.GenerateConnectToken(ctx, service.RequestGenerateConnectToken{
		GameId: responsJoinGameDTOUsers.Id,
		UserId: userId,
	})
	if err != nil {
		log.Errorf("failed to generate connect token, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(response.Join{
		Game: response.JoinGame{
			Id:           responsJoinGameDTOUsers.Id,
			Status:       responsJoinGameDTOUsers.Status,
			Name:         responsJoinGameDTOUsers.Name,
			WithPassword: responsJoinGameDTOUsers.Password != nil,
			MaxPlayers:   responsJoinGameDTOUsers.MaxPlayers,
		},
		ConnectToken: connectToken,
	})
}
