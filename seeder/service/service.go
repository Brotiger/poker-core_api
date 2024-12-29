package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/poker-core_api/seeder/config"
	"github.com/Brotiger/poker-core_api/seeder/model"
	"github.com/Brotiger/poker-core_api/seeder/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(),
	}
}

func (us *UserService) CreateUser(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(config.Cfg.MongoDB.ConnectTimeoutMs)*time.Millisecond)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.Cfg.App.Root.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	timeNow := time.Now()
	modelUser := model.User{
		Username:         config.Cfg.App.Root.Username,
		Email:            config.Cfg.App.Root.Email,
		Password:         string(hashedPassword),
		EmailConfirmed:   true,
		EmailConfirmedAt: &timeNow,
		UpdatedAt:        timeNow,
		CreatedAt:        timeNow,
	}

	if err := us.UserRepository.CreateUser(ctx, modelUser); err != nil {
		return fmt.Errorf("failed to create user, error: %w", err)
	}

	return nil
}
