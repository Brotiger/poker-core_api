package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/connection"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/model"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) FindUser(ctx context.Context, username string) (*model.User, error) {
	var modelUser model.User

	if err := connection.DB.Collection(config.Cfg.Table.User).FindOne(
		ctx,
		bson.M{"username": username},
	).Decode(&modelUser); err != nil {
		return nil, fmt.Errorf("failed to find one, error: %w", err)
	}

	return &modelUser, nil
}
