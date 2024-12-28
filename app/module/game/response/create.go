package response

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Create struct {
	Id         primitive.ObjectID   `json:"_id" example:"507f1f77bcf86cd799439011"`
	Status     string               `json:"status" example:"waiting"`
	Name       string               `json:"name" example:"test"`
	Password   *string              `json:"password,omitempty" example:"123456"`
	OwnerId    primitive.ObjectID   `json:"owner_id" example:"507f1f77bcf86cd799439011"`
	Users      []primitive.ObjectID `json:"users" example:"507f1f77bcf86cd799439011"`
	MaxPlayers int                  `json:"max_players" example:"5"`
}
