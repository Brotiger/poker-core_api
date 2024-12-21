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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (a *Auth) GenerateTokens(ctx context.Context, userId primitive.ObjectID) (*response.Login, error) {
	if err := a.Repository.DeleteRefreshToken(ctx, userId); err != nil {
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
	a.Repository.CreateRefreshToken(ctx, model.RefreshToken{
		UserId:    userId,
		Token:     refreshTokenString,
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	})

	return &response.Login{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (a *Auth) VerifyToken(tokenString string) (*TokenClaims, error) {
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
