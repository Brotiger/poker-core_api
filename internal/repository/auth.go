package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/connection"
	"github.com/Brotiger/per-painted_poker-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) FindUser(ctx context.Context, username string) (*model.User, error) {
	var modelUser model.User

	if err := connection.DB.Collection(config.Cfg.MongoDB.Collection.User).FindOne(
		ctx,
		bson.M{"username": username},
	).Decode(&modelUser); err != nil {
		return nil, fmt.Errorf("failed to find one, error: %w", err)
	}

	return &modelUser, nil
}
