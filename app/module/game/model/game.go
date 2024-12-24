package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	Id         primitive.ObjectID   `bson:"_id"`
	Status     string               `bson:"status"`
	Name       string               `bson:"name"`
	Password   *string              `bson:"password,omitempty"`
	OwnerId    primitive.ObjectID   `bson:"ownerId"`
	Users      []primitive.ObjectID `bson:"users"`
	MaxPlayers int                  `bson:"maxPlayers"`
	UpdatedAt  time.Time            `bson:"updatedAt"`
	CreatedAt  time.Time            `bson:"createdAt"`
}
