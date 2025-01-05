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
)

// @Summary Получение списка игр
// @Tags Game
// @Router /game [get]
// @Produce json
// @Param request query request.List false "Query params"
// @Success 200 {object} response.List "Успешный ответ."
// @Failure 400 {object} sharedResponse.BadRequest "Не валидный запрос."
// @Failure 401 {object} sharedResponse.Unauthorized "Невалидный токен."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (gh *GameController) List(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	requetList := request.List{
		Size: 20,
	}

	if err := c.QueryParser(&requetList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.BadRequest{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	responseGetGameListDTO, total, err := gh.gameService.GetGameList(ctx, service.RequestGetGameListDTO{
		Name: requetList.Name,
		Size: requetList.Size,
		From: requetList.From,
	})
	if err != nil {
		log.Errorf("failed to get game list, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	responseGames := []response.ListGame{}
	for _, game := range responseGetGameListDTO {
		responseGames = append(responseGames, response.ListGame{
			Id:           game.Id,
			Status:       game.Status,
			OwnerId:      game.OwnerId,
			Name:         game.Name,
			CountPlayers: game.CountPlayers,
			MaxPlayers:   game.MaxPlayers,
			WithPassword: game.WithPassword,
		})
	}

	return c.JSON(response.List{
		Total: total,
		Games: responseGames,
	})
}
