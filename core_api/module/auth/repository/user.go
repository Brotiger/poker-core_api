package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/connection"
	cError "github.com/Brotiger/poker-core_api/core_api/module/auth/error"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (ur *UserRepository) CreateUser(ctx context.Context, modelUser model.User) (*primitive.ObjectID, error) {
	result, err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).InsertOne(ctx, modelUser)
	if err != nil {
		return nil, fmt.Errorf("failed to insert one, error: %w", err)
	}

	insertId := result.InsertedID.(primitive.ObjectID)
	return &insertId, nil
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

func (ur *UserRepository) UpdateConfirmedEmailById(ctx context.Context, id primitive.ObjectID) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).UpdateOne(
		ctx,
		bson.M{
			"_id": id,
		},
		bson.M{
			"$set": bson.M{
				"emailConfirmed": true,
				"updatedAt":      time.Now(),
			},
		},
	); err != nil {
		return fmt.Errorf("failed to update one, error: %w", err)
	}

	return nil
}

func (ur *UserRepository) UpdatePasswordByUserId(ctx context.Context, id primitive.ObjectID, password string) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).UpdateOne(
		ctx,
		bson.M{
			"_id": id,
		},
		bson.M{
			"$set": bson.M{
				"password":  password,
				"updatedAt": time.Now(),
			},
		},
	); err != nil {
		return fmt.Errorf("failed to update one, error: %w", err)
	}

	return nil
}
