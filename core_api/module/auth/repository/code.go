package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/connection"
	"github.com/Brotiger/poker-core_api/core_api/module/auth/model"
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
