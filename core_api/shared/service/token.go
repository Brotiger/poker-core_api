package service

import (
	"errors"
	"fmt"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/shared/model"
	"github.com/golang-jwt/jwt"
)

var ErrInvalidToken = errors.New("invalid token")

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (ts *TokenService) VerifyToken(tokenString string) (*model.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.App.Jwt.Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse with claims, error: %w", err)
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	tokenClaims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		return nil, fmt.Errorf("failed to convert claims to model, error: %w", err)
	}

	return tokenClaims, nil
}
