package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	Id            *primitive.ObjectID `bson:"_id,omitempty"`
	UserId        primitive.ObjectID  `bson:"userId"`
	GameId        primitive.ObjectID  `bson:"gameId"`
	Status        string              `bson:"status"`
	DisconectedAt *time.Time          `bson:"disconectedAt,omitempty"`
	UpdatedAt     time.Time           `bson:"updatedAt"`
	CreatedAt     time.Time           `bson:"createdAt"`
}
