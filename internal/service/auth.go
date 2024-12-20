package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/repository"
	"github.com/Brotiger/per-painted_poker-backend/internal/request"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Repository *repository.Auth
}

func NewAuth() *Auth {
	return &Auth{
		Repository: repository.NewAuth(),
	}
}

func (a *Auth) Login(ctx context.Context, requetLogin request.Login) (*jwt.Token, error) {
	modelUser, err := a.Repository.FindUser(ctx, requetLogin.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(modelUser.Password), []byte(requetLogin.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to compare hash and password: %w", err)
	}

	claims := jwt.MapClaims{
		"id":       modelUser.Id,
		"username": modelUser.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims), nil
}
