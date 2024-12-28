package handler

import (
	"context"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/core_api/config"
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/game/request"
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/game/response"
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/game/service"
	sharedResponse "github.com/Brotiger/per-painted_poker-backend/core_api/shared/response"
	"github.com/Brotiger/per-painted_poker-backend/core_api/validator"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// @Summary Получение списка игр
// @Tags Game
// @Router /game [get]
// @Produce json
// @Param request query request.List false "Query params"
// @Success 200 {object} response.List "Успешный ответ."
// @Failure 400 {object} sharedResponse.Error400 "Не валидный запрос."
// @Failure 401 {object} sharedResponse.Error401 "Невалидный токен."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (gh *GameHandler) List(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Millisecond)
	defer cancelCtx()

	requetList := request.List{
		Size: 20,
	}

	if err := c.QueryParser(&requetList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Ошибка валидации.",
			Errors:  validator.ValidateErr(err),
		})
	}

	responseGetGameListDTO, total, err := gh.GameService.GetGameList(ctx, service.RequestGetGameListDTO{
		Name: requetList.Name,
		Size: requetList.Size,
		From: requetList.From,
	})
	if err != nil {
		log.Errorf("failed to get game list, error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	responseGames := []response.Game{}
	for _, game := range responseGetGameListDTO {
		responseGames = append(responseGames, response.Game{
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
