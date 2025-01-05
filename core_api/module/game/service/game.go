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
	GameRepository *repository.GameRepository
}

func NewGameService() *GameService {
	return &GameService{
		GameRepository: repository.NewGameRepository(),
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
	modelGames, err := gs.GameRepository.GetGames(ctx, repository.RequestGetGamesDTO{
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
			CountPlayers: len(modelGame.Users),
			MaxPlayers:   modelGame.MaxPlayers,
			WithPassword: modelGame.Password != nil,
		})
	}

	total, err := gs.GameRepository.GetGameCount(ctx, repository.RequestGetGameCountDTO{
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

	insertId, err := gs.GameRepository.CreateGame(ctx, modelGame)
	if err != nil {
		return nil, fmt.Errorf("failed to create game, error: %w", err)
	}

	return &ResponsCreateGameDTO{
		Id:         insertId,
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
	modelGame, err := gs.GameRepository.GetGameById(ctx, requestJoinGameDTO.GameId)
	if err != nil {
		return nil, fmt.Errorf("failed to get game by id, error: %w", err)
	}

	if modelGame.Password != requestJoinGameDTO.Password {
		return nil, cError.ErrComparePassword
	}

	if err := gs.GameRepository.AddUserToGame(ctx, requestJoinGameDTO.UserId, requestJoinGameDTO.GameId); err != nil {
		return nil, fmt.Errorf("failed to add user to game, error: %w", err)
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
	if err := gs.GameRepository.UpdateGameStatus(ctx, userId, "started"); err != nil {
		return fmt.Errorf("failed to update game status, error: %w", err)
	}

	return nil
}

func (gs *GameService) UserHasGame(ctx context.Context, userId primitive.ObjectID) (bool, error) {
	count, err := gs.GameRepository.CountUserGames(ctx, userId)
	if err != nil {
		return false, fmt.Errorf("failed to count user games count, error: %w", err)
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (gs *GameService) GameCanBeStarted(ctx context.Context, userId primitive.ObjectID) (bool, error) {
	modelGame, err := gs.GameRepository.GetGameByOwnerId(ctx, userId)
	if err != nil {
		if err == cError.ErrGameNotFound {
			return false, nil
		}

		return false, fmt.Errorf("failed to get game by owner id, error: %w", err)
	}

	if modelGame.Status != "waiting" {
		return false, nil
	}

	if modelGame.MaxPlayers != len(modelGame.Users) {
		return false, nil
	}

	return true, nil
}
