package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/per-painted_poker-backend/app/config"
	"github.com/Brotiger/per-painted_poker-backend/app/connection"
	cError "github.com/Brotiger/per-painted_poker-backend/app/module/auth/error"
	"github.com/Brotiger/per-painted_poker-backend/app/module/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		options.FindOne().SetHint(bson.D{
			{Key: "email", Value: 1},
		}),
	).Decode(&modelUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, cError.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find one, error: %w", err)
	}

	return &modelUser, nil
}
