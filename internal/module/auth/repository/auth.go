package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/connection"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) FindUser(ctx context.Context, username string) (*model.User, error) {
	var modelUser model.User

	if err := connection.DB.Collection(config.Cfg.Table.User).FindOne(
		ctx,
		bson.M{"username": username},
	).Decode(&modelUser); err != nil {
		return nil, fmt.Errorf("failed to find one, error: %w", err)
	}

	return &modelUser, nil
}

func (a *Auth) CreateRefreshToken(ctx context.Context, modelRefreshToken model.RefreshToken) error {
	if _, err := connection.DB.Collection(config.Cfg.Table.RefreshToken).InsertOne(
		ctx,
		modelRefreshToken,
	); err != nil {
		return fmt.Errorf("failed to insert one, error: %w", err)
	}

	return nil
}

func (a *Auth) DeleteRefreshToken(ctx context.Context, userId primitive.ObjectID) error {
	if _, err := connection.DB.Collection(config.Cfg.Table.RefreshToken).DeleteOne(
		ctx,
		bson.M{"userId": userId},
	); err != nil {
		return fmt.Errorf("failed to delete one, error: %w", err)
	}

	return nil
}
