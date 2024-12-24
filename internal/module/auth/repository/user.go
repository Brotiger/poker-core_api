package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/Brotiger/per-painted_poker-backend/internal/connection"
	cError "github.com/Brotiger/per-painted_poker-backend/internal/module/auth/error"
	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var modelUser model.User
	if err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).FindOne(
		ctx,
		bson.M{"email": email},
	).Decode(&modelUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, cError.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find one, error: %w", err)
	}

	return &modelUser, nil
}
