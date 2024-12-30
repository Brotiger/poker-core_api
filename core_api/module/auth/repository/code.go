package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/connection"
	cError "github.com/Brotiger/poker-core_api/core_api/module/auth/error"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CodeRepository struct{}

func NewCodeRepository() *CodeRepository {
	return &CodeRepository{}
}

func (cr *CodeRepository) CreateCode(ctx context.Context, modelCode model.Code) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Code).InsertOne(ctx, modelCode); err != nil {
		return fmt.Errorf("failed to insert one, error: %w", err)
	}

	return nil
}

func (cr *CodeRepository) FindCodeByUserId(ctx context.Context, id primitive.ObjectID) (*model.Code, error) {
	var modelCode model.Code
	if err := connection.DB.Collection(config.Cfg.MongoDB.Table.Code).FindOne(
		ctx,
		bson.M{
			"userId": id,
		},
	).Decode(&modelCode); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, cError.ErrCodeNotFound
		}

		return nil, fmt.Errorf("failed to find one, error: %w", err)
	}

	return &modelCode, nil
}

func (ct *CodeRepository) DeleteById(ctx context.Context, id primitive.ObjectID) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.Code).DeleteOne(
		ctx,
		bson.M{
			"_id": id,
		},
	); err != nil {
		return fmt.Errorf("failed to delete one, error: %w", err)
	}

	return nil
}
