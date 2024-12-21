package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/model"
	"github.com/Brotiger/per-painted_poker-backend/internal/repository"
	"github.com/Brotiger/per-painted_poker-backend/internal/request"
	"github.com/Brotiger/per-painted_poker-backend/internal/response"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Repository *repository.Auth
}

type TokenClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func NewAuth() *Auth {
	return &Auth{
		Repository: repository.NewAuth(),
	}
}

func (a *Auth) Login(ctx context.Context, requetLogin request.Login) (*model.User, error) {
	modelUser, err := a.Repository.FindUser(ctx, requetLogin.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to find user, error: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(modelUser.Password), []byte(requetLogin.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to compare hash and password, error: %w", err)
	}

	return modelUser, nil
}

func (a *Auth) GenerateTokens(ctx context.Context, modelUser model.User) (*response.Login, error) {
	if err := a.Repository.DeleteRefreshToken(ctx, modelUser.Id); err != nil {
		return nil, fmt.Errorf("failed to delete refresh token, error: %w", err)
	}

	accessTokenClaims := TokenClaims{
		UserId: modelUser.Id.Hex(),
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
		UserId: modelUser.Id.Hex(),
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
	a.Repository.CreateRefreshToken(ctx, model.RefreshToken{
		UserId:    modelUser.Id,
		Token:     refreshTokenString,
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	})

	return &response.Login{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
