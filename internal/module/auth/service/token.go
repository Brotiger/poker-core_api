package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/model"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/repository"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/response"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshToken struct {
	RefreshTokenRepository *repository.RefreshToken
}

type TokenClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func NewRefreshToken() *RefreshToken {
	return &RefreshToken{
		RefreshTokenRepository: repository.NewRefreshToken(),
	}
}

func (rt *RefreshToken) GenerateTokens(ctx context.Context, userId primitive.ObjectID) (*response.Token, error) {
	if err := rt.RefreshTokenRepository.DeleteRefreshToken(ctx, userId); err != nil {
		return nil, fmt.Errorf("failed to delete refresh token, error: %w", err)
	}

	accessTokenClaims := TokenClaims{
		UserId: userId.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Cfg.App.Jwt.AccessTokenExpireAt) * time.Minute).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(config.Cfg.App.Jwt.Secret))
	if err != nil {
		return nil, fmt.Errorf("failed to signed string, error: %w", err)
	}

	refreshTokenClaims := TokenClaims{
		UserId: userId.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Cfg.App.Jwt.RefreshTokenExpireAt) * time.Minute).Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(config.Cfg.App.Jwt.Secret))
	if err != nil {
		return nil, fmt.Errorf("failed to signed string, error %w", err)
	}

	timeNow := time.Now()
	rt.RefreshTokenRepository.CreateRefreshToken(ctx, model.RefreshToken{
		UserId:    userId,
		Token:     refreshTokenString,
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	})

	return &response.Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (rt *RefreshToken) VerifyToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.App.Jwt.Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse with claims, error: %w", err)
	}

	tokenClaims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return tokenClaims, nil
}
