package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	CreatedAt time.Time          `bson:"createdAt"`
}
