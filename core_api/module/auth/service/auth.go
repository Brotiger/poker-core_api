package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/connection"
	cError "github.com/Brotiger/poker-core_api/core_api/module/auth/error"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/model"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/repository"
	sharedService "github.com/Brotiger/poker-core_api/core_api/shared/service"
	natsModel "github.com/Brotiger/poker-core_api/pkg/nats/model"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repository.UserRepository
	codeRepository *repository.CodeRepository
	randomService  *sharedService.RandomService
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepository: repository.NewUserRepository(),
		codeRepository: repository.NewCodeRepository(),
		randomService:  sharedService.NewRandomService(),
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
	modelUser, err := as.userRepository.FindUserByEmail(ctx, requestGetUserDTO.Email)
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
		Id: *modelUser.Id,
	}, nil
}

type RequestRegisterDTO struct {
	Email    string
	Username string
	Password string
}

func (as *AuthService) Register(ctx context.Context, requestRegisterDTO RequestRegisterDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestRegisterDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	timeNow := time.Now()
	modelUser := model.User{
		Username:  requestRegisterDTO.Username,
		Email:     requestRegisterDTO.Email,
		Password:  string(hashedPassword),
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}

	userId, err := as.userRepository.CreateUser(ctx, modelUser)
	if err != nil {
		return fmt.Errorf("failed to create user, error: %w", err)
	}

	codeType := "register"
	modelCode := model.Code{
		UserId:    *userId,
		Code:      as.randomService.RandomString(config.Cfg.App.CodeLength),
		Type:      codeType,
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}

	if err := as.codeRepository.CreateCode(ctx, modelCode); err != nil {
		return fmt.Errorf("failed to create code, error: %w", err)
	}

	data, err := json.Marshal(natsModel.Register{
		Code: modelCode.Code,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal mailer msg, error: %w", err)
	}

	if _, err := connection.JS.PublishMsg(ctx, &nats.Msg{
		Subject: config.Cfg.Nats.Streams.Mailer,
		Header: nats.Header{
			"type":  []string{codeType},
			"email": []string{requestRegisterDTO.Email},
		},
		Data: data,
	}); err != nil {
		return fmt.Errorf("failed to publish nats msg, error: %w", err)
	}

	return nil
}

func (as *AuthService) CheckUsername(ctx context.Context, username string) (bool, error) {
	count, err := as.userRepository.CountUsersByUsername(ctx, username)
	if err != nil {
		return false, fmt.Errorf("failed to count users, error: %w", err)
	}

	if count > 0 {
		return false, nil
	}

	return true, nil
}

func (as *AuthService) CheckEmail(ctx context.Context, email string) (bool, error) {
	count, err := as.userRepository.CountUsersByEmail(ctx, email)
	if err != nil {
		return false, fmt.Errorf("failed to count users, error: %w", err)
	}

	if count > 0 {
		return false, nil
	}

	return true, nil
}

type RequestConfirmedEmailDTO struct {
	Email string
	Code  string
}

func (as *AuthService) ConfirmEmail(ctx context.Context, requestConfirmedEmailDTO RequestConfirmedEmailDTO) error {
	modelUser, err := as.userRepository.FindUserByEmail(ctx, requestConfirmedEmailDTO.Email)
	if err != nil {
		return fmt.Errorf("failed to find user by email, error: %w", err)
	}

	modelCode, err := as.codeRepository.FindCodeByUserId(ctx, *modelUser.Id)
	if err != nil {
		return fmt.Errorf("failed to find code, error: %w", err)
	}

	if modelCode.Code != requestConfirmedEmailDTO.Code {
		return cError.ErrCompareCode
	}

	if err := as.codeRepository.DeleteById(ctx, *modelCode.Id); err != nil {
		return fmt.Errorf("failed to delete code by id, error: %w", err)
	}

	if err := as.userRepository.UpdateConfirmedEmailById(ctx, *modelUser.Id); err != nil {
		return fmt.Errorf("failed to update user by id, error: %w", err)
	}

	return nil
}

func (as *AuthService) Restore(ctx context.Context, email string) error {
	modelUser, err := as.userRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to find user, error: %w", err)
	}

	timeNow := time.Now()
	codeType := "restore"
	modelCode := model.Code{
		UserId:    *modelUser.Id,
		Code:      as.randomService.RandomString(config.Cfg.App.CodeLength),
		Type:      codeType,
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}

	if err := as.codeRepository.CreateCode(ctx, modelCode); err != nil {
		return fmt.Errorf("failed to create code, error: %w", err)
	}

	data, err := json.Marshal(natsModel.Restore{
		Code: modelCode.Code,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal mailer msg, error: %w", err)
	}

	if _, err := connection.JS.PublishMsg(ctx, &nats.Msg{
		Subject: config.Cfg.Nats.Streams.Mailer,
		Header: nats.Header{
			"type":  []string{codeType},
			"email": []string{email},
		},
		Data: data,
	}); err != nil {
		return fmt.Errorf("failed to publish nats msg, error: %w", err)
	}

	return nil
}

type RequestConfirmedRestoreDTO struct {
	Email    string
	Password string
	Code     string
}

func (as *AuthService) ConfirmRestore(ctx context.Context, requestConfirmedRestoreDTO RequestConfirmedRestoreDTO) error {
	modelUser, err := as.userRepository.FindUserByEmail(ctx, requestConfirmedRestoreDTO.Email)
	if err != nil {
		return fmt.Errorf("failed to find user by email, error: %w", err)
	}

	modelCode, err := as.codeRepository.FindCodeByUserId(ctx, *modelUser.Id)
	if err != nil {
		return fmt.Errorf("failed to find code, error: %w", err)
	}

	if modelCode.Code != requestConfirmedRestoreDTO.Code {
		return cError.ErrCompareCode
	}

	if err := as.codeRepository.DeleteById(ctx, *modelCode.Id); err != nil {
		return fmt.Errorf("failed to delete code by id, error: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestConfirmedRestoreDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	if err := as.userRepository.UpdatePasswordByUserId(ctx, *modelUser.Id, string(hashedPassword)); err != nil {
		return fmt.Errorf("failed to update user by id, error: %w", err)
	}

	return nil
}
