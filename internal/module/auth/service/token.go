package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/dto"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/model"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/repository"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	sharedModel "github.com/Brotiger/per-painted_poker-backend/internal/shared/model"
)

type RefreshToken struct {
	RefreshTokenRepository *repository.RefreshToken
}

func NewRefreshToken() *RefreshToken {
	return &RefreshToken{
		RefreshTokenRepository: repository.NewRefreshToken(),
	}
}

func (rt *RefreshToken) GenerateTokens(ctx context.Context, userId primitive.ObjectID) (*dto.Token, error) {
	if err := rt.RefreshTokenRepository.DeleteRefreshToken(ctx, userId); err != nil {
		return nil, fmt.Errorf("failed to delete refresh token, error: %w", err)
	}

	accessTokenClaims := sharedModel.TokenClaims{
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

	refreshTokenClaims := sharedModel.TokenClaims{
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

	return &dto.Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (rt *RefreshToken) CheckRefreshToken(ctx context.Context, userId primitive.ObjectID) (bool, error) {
	count, err := rt.RefreshTokenRepository.CountRefreshToken(ctx, userId)
	if err != nil {
		return false, fmt.Errorf("failed to count refresh token, error: %w", err)
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (rt *RefreshToken) Logout(ctx context.Context, userId primitive.ObjectID) error {
	if err := rt.RefreshTokenRepository.DeleteRefreshToken(ctx, userId); err != nil {
		return fmt.Errorf("failed to delete refresh token, error: %w", err)
	}

	return nil
}
