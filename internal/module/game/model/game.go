package model

import (
	"time"

	"github.com/Brotiger/per-painted_poker-backend/internal/module/auth/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	Password   *string            `bson:"password,omitempty"`
	Users      []model.User       `bson:"users"`
	MaxPlayers int                `bson:"maxPlayers"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
	CreatedAt  time.Time          `bson:"createdAt"`
}
