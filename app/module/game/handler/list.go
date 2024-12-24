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
func (a *Game) List(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Microsecond)
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

	modelGames, total, err := a.GameService.GetGameList(ctx, requetList)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var responseGames []response.Game
	for _, modelGame := range modelGames {
		responseGames = append(responseGames, response.Game{
			Id:           modelGame.Id,
			Status:       modelGame.Status,
			OwnerId:      modelGame.OwnerId,
			Name:         modelGame.Name,
			CountPlayers: len(modelGame.Users),
			MaxPlayers:   modelGame.MaxPlayers,
			WithPassword: modelGame.Password != nil,
		})
	}

	return c.JSON(response.List{
		Total: total,
		Games: responseGames,
	})
}
