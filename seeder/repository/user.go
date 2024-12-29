package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/poker-core_api/seeder/config"
	"github.com/Brotiger/poker-core_api/seeder/connection"
	"github.com/Brotiger/poker-core_api/seeder/model"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) CreateUser(ctx context.Context, modelUser model.User) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).InsertOne(ctx, modelUser); err != nil {
		return fmt.Errorf("failed to insert one, error: %w", err)
	}

	return nil
}
