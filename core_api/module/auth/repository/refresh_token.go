package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/core_api/config"
	"github.com/Brotiger/per-painted_poker-backend/core_api/connection"
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RefreshTokenRepository struct{}

func NewRefreshTokenRepository() *RefreshTokenRepository {
	return &RefreshTokenRepository{}
}

func (rt *RefreshTokenRepository) CreateRefreshToken(ctx context.Context, modelRefreshToken model.RefreshToken) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.RefreshToken).InsertOne(
		ctx,
		modelRefreshToken,
	); err != nil {
		return fmt.Errorf("failed to insert one, error: %w", err)
	}

	return nil
}

func (rt *RefreshTokenRepository) DeleteRefreshToken(ctx context.Context, userId primitive.ObjectID) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.RefreshToken).DeleteOne(
		ctx,
		bson.M{"userId": userId},
		options.Delete().SetHint(bson.D{
			{Key: "userId", Value: 1},
		}),
	); err != nil {
		return fmt.Errorf("failed to delete one, error: %w", err)
	}

	return nil
}

func (rt *RefreshTokenRepository) CountRefreshToken(ctx context.Context, token string) (int64, error) {
	count, err := connection.DB.Collection(config.Cfg.MongoDB.Table.RefreshToken).CountDocuments(
		ctx,
		bson.M{
			"token": token,
		},
		options.Count().SetHint(bson.D{
			{Key: "token", Value: 1},
		}),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to count documents, error: %w", err)
	}

	return count, nil
}
