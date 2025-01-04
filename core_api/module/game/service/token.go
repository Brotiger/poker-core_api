package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	pkgModel "github.com/Brotiger/poker-core_api/pkg/model"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (ts *TokenService) GenerateConnectToken(ctx context.Context, gameId primitive.ObjectID) (string, error) {
	connectTokenClaims := pkgModel.ConnectClaims{
		GameId: gameId.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Cfg.App.ConnectTokenExpireAt) * time.Millisecond).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, connectTokenClaims)
	connectTokenString, err := accessToken.SignedString([]byte(config.Cfg.App.Jwt.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to signed string, error: %w", err)
	}

	return connectTokenString, nil
}
