package service

import (
	"context"
	"fmt"
	"time"

	cError "github.com/Brotiger/per-painted_poker-backend/core_api/module/auth/error"
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/auth/model"
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/auth/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		UserRepository: repository.NewUserRepository(),
	}
}

type RequestGetUserDTO struct {
	Email    string
	Password string
}

type ResponseGetUserDTO struct {
	Id primitive.ObjectID
}

func (as *AuthService) GetUser(ctx context.Context, requestGetUserDTO RequestGetUserDTO) (*ResponseGetUserDTO, error) {
	modelUser, err := as.UserRepository.FindUserByEmail(ctx, requestGetUserDTO.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user, error: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(modelUser.Password), []byte(requestGetUserDTO.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, cError.ErrCompareHashAndPassword
		}

		return nil, fmt.Errorf("failed to compare hash and password, error: %w", err)
	}

	return &ResponseGetUserDTO{
		Id: modelUser.Id,
	}, nil
}

type RequestRegisterDTO struct {
	Email    string
	Username string
	Password string
}

func (as *AuthService) Register(ctx context.Context, requetRegisterDTO RequestRegisterDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requetRegisterDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	timeNow := time.Now()
	modelUser := model.User{
		Username: requetRegisterDTO.Username,
		Email:    requetRegisterDTO.Email,
		Password: string(hashedPassword),

		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}

	if err := as.UserRepository.CreateUser(ctx, modelUser); err != nil {
		return fmt.Errorf("failed to create user, error: %w", err)
	}

	return nil
}
