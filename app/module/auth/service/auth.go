package service

import (
	"context"
	"fmt"

	cError "github.com/Brotiger/per-painted_poker-backend/app/module/auth/error"
	"github.com/Brotiger/per-painted_poker-backend/app/module/auth/model"
	"github.com/Brotiger/per-painted_poker-backend/app/module/auth/repository"
	"github.com/Brotiger/per-painted_poker-backend/app/module/auth/request"
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

func (a *Auth) GetUser(ctx context.Context, requetLogin request.Login) (*model.User, error) {
	modelUser, err := a.UserRepository.FindUserByEmail(ctx, requetLogin.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user, error: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(modelUser.Password), []byte(requetLogin.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, cError.ErrCompareHashAndPassword
		}

		return nil, fmt.Errorf("failed to compare hash and password, error: %w", err)
	}

	return modelUser, nil
}
