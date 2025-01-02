package response

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Create struct {
	Game         CreateGame `json:"game"`
	ConnectToken string     `json:"connect_token"`
}

type CreateGame struct {
	Id           primitive.ObjectID   `json:"_id" example:"507f1f77bcf86cd799439011"`
	Status       string               `json:"status" example:"waiting"`
	Name         string               `json:"name" example:"test"`
	WithPassword bool                 `json:"with_password"`
	Users        []primitive.ObjectID `json:"users" example:"507f1f77bcf86cd799439011"`
	MaxPlayers   int                  `json:"max_players" example:"5"`
}
