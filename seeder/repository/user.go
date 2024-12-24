package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/seeder/config"
	"github.com/Brotiger/per-painted_poker-backend/seeder/connection"
	"github.com/Brotiger/per-painted_poker-backend/seeder/model"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) CreateUser(ctx context.Context, modelUser model.User) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).InsertOne(ctx, modelUser); err != nil {
		return fmt.Errorf("failed to insert one, error: %w", err)
	}

	return nil
}
