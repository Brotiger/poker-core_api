package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/poker-core_api/core_api/config"
	"github.com/Brotiger/poker-core_api/core_api/connection"
	"github.com/Brotiger/poker-core_api/core_api/module/game/model"
)

type ConnectTokenRepository struct{}

func NewConnectTokenRepository() *ConnectTokenRepository {
	return &ConnectTokenRepository{}
}

func (ctr *ConnectTokenRepository) CreateJoinToken(ctx context.Context, modelJoinToken model.ConnectToken) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.JoinToken).InsertOne(
		ctx,
		modelJoinToken,
	); err != nil {
		return fmt.Errorf("failed to insert one, error: %w", err)
	}

	return nil
}
