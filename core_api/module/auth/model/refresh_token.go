package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshToken struct {
	Id        *primitive.ObjectID `bson:"_id,omitempty"`
	UserId    primitive.ObjectID  `bson:"userId"`
	Token     string              `bson:"token"`
	UpdatedAt time.Time           `bson:"updatedAt"`
	CreatedAt time.Time           `bson:"createdAt"`
}
