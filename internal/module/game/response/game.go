package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	Id           primitive.ObjectID `json:"id"`
	Name         string             `json:"title"`
	CountPlayers int                `json:"count_players"`
	MaxPlayers   int                `json:"max_players"`
	WithPassword bool               `json:"with_password"`
}
