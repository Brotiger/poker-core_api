package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/pkg/model"
	"github.com/golang-jwt/jwt"
)

const headerPrefix = "Bearer"

var ErrInvalidToken = errors.New("invalid token")

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (ts *TokenService) VerifyToken(tokenString string) (*model.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.App.Jwt.Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse with claims, error: %w", err)
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	tokenClaims, ok := token.Claims.(*model.JWTClaims)
	if !ok {
		return nil, fmt.Errorf("failed to convert claims to model, error: %w", err)
	}

	return tokenClaims, nil
}

func (ts *TokenService) GetToken(accessToken string) (string, error) {
	l := len(headerPrefix)
	if len(accessToken) < l+2 || accessToken[:l] != headerPrefix {
		return "", errors.New("invalid token format")
	}

	return strings.TrimSpace(accessToken[l:]), nil
}
