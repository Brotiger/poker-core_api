package service

import (
	"context"
	"fmt"
	"time"

	cError "github.com/Brotiger/poker-core_api/core_api/module/game/error"
	"github.com/Brotiger/poker-core_api/core_api/module/game/model"
	"github.com/Brotiger/poker-core_api/core_api/module/game/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameService struct {
	gameRepository   *repository.GameRepository
	playerRepository *repository.PlayerRepository
}

func NewGameService() *GameService {
	return &GameService{
		gameRepository:   repository.NewGameRepository(),
		playerRepository: repository.NewPlayerRepository(),
	}
}

type RequestGetGameListDTO struct {
	Name *string
	From int64
	Size int64
}

type ResponsGetGameListDTO struct {
	Id           primitive.ObjectID
	Status       string
	OwnerId      primitive.ObjectID
	Name         string
	CountPlayers int
	MaxPlayers   int
	WithPassword bool
}

func (gs *GameService) GetGameList(ctx context.Context, requestGetGameListDTO RequestGetGameListDTO) ([]ResponsGetGameListDTO, int64, error) {
	modelGames, err := gs.gameRepository.GetGames(ctx, repository.RequestGetGamesDTO{
		Name: requestGetGameListDTO.Name,
		From: requestGetGameListDTO.From,
		Size: requestGetGameListDTO.Size,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get games, error: %w", err)
	}

	var responsGetGameListDTO []ResponsGetGameListDTO
	for _, modelGame := range modelGames {
		responsGetGameListDTO = append(responsGetGameListDTO, ResponsGetGameListDTO{
			Id:           *modelGame.Id,
			Status:       modelGame.Status,
			OwnerId:      modelGame.OwnerId,
			Name:         modelGame.Name,
			CountPlayers: modelGame.CountPlayers,
			MaxPlayers:   modelGame.MaxPlayers,
			WithPassword: modelGame.Password != nil,
		})
	}

	total, err := gs.gameRepository.GetGameCount(ctx, repository.RequestGetGameCountDTO{
		Name: requestGetGameListDTO.Name,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get game count, error: %w", err)
	}

	return responsGetGameListDTO, total, nil
}

type RequestCreateGameDTO struct {
	UserId     primitive.ObjectID
	Name       string
	MaxPlayers int
	Password   *string
}

type ResponsCreateGameDTO struct {
	Id         primitive.ObjectID
	Status     string
	Name       string
	Password   *string
	MaxPlayers int
}

func (gs *GameService) CreateGame(ctx context.Context, requestCreateGameDTO RequestCreateGameDTO) (*ResponsCreateGameDTO, error) {
	timeNow := time.Now()
	modelGame := model.Game{
		Status:     "waiting",
		Name:       requestCreateGameDTO.Name,
		Password:   requestCreateGameDTO.Password,
		MaxPlayers: requestCreateGameDTO.MaxPlayers,
		OwnerId:    requestCreateGameDTO.UserId,
		UpdatedAt:  timeNow,
		CreatedAt:  timeNow,
	}

	insertId, err := gs.gameRepository.CreateGame(ctx, modelGame)
	if err != nil {
		return nil, fmt.Errorf("failed to create game, error: %w", err)
	}

	if _, err := gs.playerRepository.CreatePlayer(ctx, model.Player{
		GameId:    *insertId,
		UserId:    requestCreateGameDTO.UserId,
		Status:    "waiting",
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}); err != nil {
		return nil, fmt.Errorf("failed to create player, error: %w", err)
	}

	if err := gs.gameRepository.IncCountPlayers(ctx, *insertId); err != nil {
		return nil, fmt.Errorf("failed to inc count players, error: %w", err)
	}

	return &ResponsCreateGameDTO{
		Id:         *insertId,
		Status:     "waiting",
		Name:       requestCreateGameDTO.Name,
		Password:   requestCreateGameDTO.Password,
		MaxPlayers: requestCreateGameDTO.MaxPlayers,
	}, nil
}

type RequestJoinGameDTO struct {
	GameId   primitive.ObjectID
	UserId   primitive.ObjectID
	Password *string
}

type ResponsJoinGameDTO struct {
	Id         primitive.ObjectID
	Status     string
	Name       string
	Password   *string
	MaxPlayers int
}

func (gs *GameService) JoinGame(ctx context.Context, requestJoinGameDTO RequestJoinGameDTO) (*ResponsJoinGameDTO, error) {
	modelGame, err := gs.gameRepository.GetGameById(ctx, requestJoinGameDTO.GameId)
	if err != nil {
		return nil, fmt.Errorf("failed to get game by id, error: %w", err)
	}

	if modelGame.Password != requestJoinGameDTO.Password {
		return nil, cError.ErrComparePassword
	}

	timeNow := time.Now()
	if _, err := gs.playerRepository.CreatePlayer(ctx, model.Player{
		GameId:    requestJoinGameDTO.GameId,
		UserId:    requestJoinGameDTO.UserId,
		Status:    "waiting",
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}); err != nil {
		return nil, fmt.Errorf("failed to create player, error: %w", err)
	}

	if err := gs.gameRepository.IncCountPlayers(ctx, requestJoinGameDTO.GameId); err != nil {
		return nil, fmt.Errorf("failed to inc count players, error: %w", err)
	}

	return &ResponsJoinGameDTO{
		Id:         *modelGame.Id,
		Status:     modelGame.Status,
		Name:       modelGame.Name,
		Password:   modelGame.Password,
		MaxPlayers: modelGame.MaxPlayers,
	}, nil
}

func (gs *GameService) StartGame(ctx context.Context, userId primitive.ObjectID) error {
	if err := gs.gameRepository.UpdateGameStatus(ctx, userId, "started"); err != nil {
		return fmt.Errorf("failed to update game status, error: %w", err)
	}

	return nil
}

func (gs *GameService) UserHasGame(ctx context.Context, userId primitive.ObjectID) (bool, error) {
	count, err := gs.gameRepository.CountUserGames(ctx, userId)
	if err != nil {
		return false, fmt.Errorf("failed to count user games count, error: %w", err)
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (gs *GameService) CheckGameAllowToJoin(ctx context.Context, gameId primitive.ObjectID) (bool, error) {
	modelGame, err := gs.gameRepository.GetGameById(ctx, gameId)
	if err != nil {
		return false, fmt.Errorf("failed to get game by id, error: %w", err)
	}

	if modelGame.Status != "waiting" {
		return false, nil
	}

	countPlayers, err := gs.playerRepository.CountPlayersByGameId(ctx, gameId)
	if err != nil {
		return false, fmt.Errorf("failed to count players by game id, error: %w", err)
	}

	if int(countPlayers) >= modelGame.MaxPlayers {
		return false, nil
	}

	return true, nil
}
