package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/core_api/config"
	"github.com/Brotiger/per-painted_poker-backend/core_api/connection"
	cError "github.com/Brotiger/per-painted_poker-backend/core_api/module/auth/error"
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var modelUser model.User
	if err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).FindOne(
		ctx,
		bson.M{
			"email":          email,
			"emailConfirmed": true,
		},
		options.FindOne().SetHint(bson.D{
			{Key: "email", Value: 1},
			{Key: "emailConfirmed", Value: 1},
		}),
	).Decode(&modelUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, cError.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find one, error: %w", err)
	}

	return &modelUser, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, modelUser model.User) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).InsertOne(ctx, modelUser); err != nil {
		return fmt.Errorf("failed to insert one, error: %w", err)
	}

	return nil
}

func (ur *UserRepository) CountUsersByUsername(ctx context.Context, username string) (int64, error) {
	count, err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).CountDocuments(
		ctx,
		bson.M{
			"username": username,
		},
		options.Count().SetHint(bson.D{
			{Key: "username", Value: 1},
		}),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to count documents, error: %w", err)
	}

	return count, nil
}

func (ur *UserRepository) CountUsersByEmail(ctx context.Context, email string) (int64, error) {
	count, err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).CountDocuments(
		ctx,
		bson.M{
			"email": email,
		},
		options.Count().SetHint(bson.D{
			{Key: "email", Value: 1},
		}),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to count documents, error: %w", err)
	}

	return count, nil
}
