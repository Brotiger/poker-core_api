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
)

// @Summary Игра
// @Tags Game
// @Router /game/list [get]
// @Produce json
// @Success 200 {object} response.GameList "Успешный ответ."
// @Failure 400 {object} sharedResponse.Error400 "Не валидный запрос."
// @Failure 401 {object} sharedResponse.Error401 "Не верное имя пользователя или пароль."
// @Failure 500 "Ошибка сервера."
// @securityDefinitions.apikey Authorization
// @in header
// @Security Authorization
func (a *Game) List(c *fiber.Ctx) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(config.Cfg.Fiber.RequestTimeoutMs)*time.Microsecond)
	defer cancelCtx()

	var requetList request.List
	if err := c.QueryParser(&requetList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Не валидный запрос.",
		})
	}

	if err := validator.Validator.Struct(requetList); err != nil {
		fieldErrors := validator.ValidateErr(err)

		return c.Status(fiber.StatusBadRequest).JSON(sharedResponse.Error400{
			Message: "Ошибка валидации.",
			Errors:  fieldErrors,
		})
	}

	modelGames, total, err := a.GameService.GetList(ctx, requetList)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var responseGames []response.Game
	for _, modelGame := range modelGames {
		responseGames = append(responseGames, response.Game{
			Id:           modelGame.Id,
			Name:         modelGame.Name,
			CountPlayers: len(modelGame.Users),
			MaxPlayers:   modelGame.MaxPlayers,
			WithPassword: modelGame.Password != nil,
		})
	}

	return c.JSON(response.GameList{
		Total: total,
		Games: responseGames,
	})
}
