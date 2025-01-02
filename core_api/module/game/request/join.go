package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type Join struct {
	GameId   primitive.ObjectID `json:"gameId" example:"507f1f77bcf86cd799439011"`
	Password *string            `json:"password,omitempty" example:"123456"`
}
