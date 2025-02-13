package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Code struct {
	Id        *primitive.ObjectID `bson:"_id,omitempty"`
	UserId    primitive.ObjectID  `bson:"userId"`
	Code      string              `bson:"code"`
	Type      string              `bson:"type"`
	UpdatedAt time.Time           `bson:"updatedAt"`
	CreatedAt time.Time           `bson:"createdAt"`
}
