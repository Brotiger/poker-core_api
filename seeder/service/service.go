package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/per-painted_poker-backend/seeder/config"
	"github.com/Brotiger/per-painted_poker-backend/seeder/model"
	"github.com/Brotiger/per-painted_poker-backend/seeder/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserRepository *repository.User
}

func NewUser() *User {
	return &User{
		UserRepository: repository.NewUser(),
	}
}

func (u *User) CreateUser(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(config.Cfg.MongoDB.ConnectTimeoutMs)*time.Millisecond)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.Cfg.App.Root.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	timeNow := time.Now()
	modelUser := model.User{
		Username:  config.Cfg.App.Root.Username,
		Email:     config.Cfg.App.Root.Email,
		Password:  string(hashedPassword),
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}

	if err := u.UserRepository.CreateUser(ctx, modelUser); err != nil {
		return fmt.Errorf("failed to create user, error: %w", err)
	}

	return nil
}
