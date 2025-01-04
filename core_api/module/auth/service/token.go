package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/model"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/repository"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	pkgModel "github.com/Brotiger/poker-core_api/pkg/model"
)

type RefreshTokenService struct {
	RefreshTokenRepository *repository.RefreshTokenRepository
}

func NewRefreshTokenService() *RefreshTokenService {
	return &RefreshTokenService{
		RefreshTokenRepository: repository.NewRefreshTokenRepository(),
	}
}

type ResponseTokenDTO struct {
	AccessToken  string
	RefreshToken string
}

func (rt *RefreshTokenService) GenerateTokens(ctx context.Context, userId primitive.ObjectID) (*ResponseTokenDTO, error) {
	if err := rt.RefreshTokenRepository.DeleteRefreshToken(ctx, userId); err != nil {
		return nil, fmt.Errorf("failed to delete refresh token, error: %w", err)
	}

	accessTokenClaims := pkgModel.JWTClaims{
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

	refreshTokenClaims := pkgModel.JWTClaims{
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

	return &ResponseTokenDTO{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (rt *RefreshTokenService) CheckRefreshToken(ctx context.Context, token string) (bool, error) {
	count, err := rt.RefreshTokenRepository.CountRefreshToken(ctx, token)
	if err != nil {
		return false, fmt.Errorf("failed to count refresh token, error: %w", err)
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (rt *RefreshTokenService) Logout(ctx context.Context, userId primitive.ObjectID) error {
	if err := rt.RefreshTokenRepository.DeleteRefreshToken(ctx, userId); err != nil {
		return fmt.Errorf("failed to delete refresh token, error: %w", err)
	}

	return nil
}
