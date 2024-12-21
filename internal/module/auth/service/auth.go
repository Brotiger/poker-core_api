package service

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/model"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/repository"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/request"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	UserRepository *repository.User
}

func NewAuth() *Auth {
	return &Auth{
		UserRepository: repository.NewUser(),
	}
}

func (a *Auth) Login(ctx context.Context, requetLogin request.Login) (*model.User, error) {
	modelUser, err := a.UserRepository.FindUser(ctx, requetLogin.Username)
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
