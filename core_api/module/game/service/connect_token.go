package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/module/game/model"
	"github.com/Brotiger/poker-core_api/core_api/module/game/repository"
	sharedService "github.com/Brotiger/poker-core_api/core_api/shared/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectTokenService struct {
	connectTokenRepository *repository.ConnectTokenRepository
	randomService          *sharedService.RandomService
}

func NewConnectTokenService() *ConnectTokenService {
	return &ConnectTokenService{
		connectTokenRepository: repository.NewConnectTokenRepository(),
		randomService:          sharedService.NewRandomService(),
	}
}

type RequestGenerateConnectToken struct {
	GameId primitive.ObjectID
	UserId primitive.ObjectID
}

func (cs *ConnectTokenService) GenerateConnectToken(ctx context.Context, requestGenerateConnectToken RequestGenerateConnectToken) (string, error) {
	token := cs.randomService.RandomString(config.Cfg.ConnectToken.Length)
	timeNow := time.Now()
	if err := cs.connectTokenRepository.CreateJoinToken(ctx, model.ConnectToken{
		Token:     token,
		UserId:    requestGenerateConnectToken.UserId,
		GameId:    requestGenerateConnectToken.GameId,
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}); err != nil {
		return "", fmt.Errorf("failed to create join token, error: %w", err)
	}

	return token, nil
}
