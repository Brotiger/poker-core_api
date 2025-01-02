package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type List struct {
	Total int64      `json:"total" example:"100"`
	Games []ListGame `json:"games"`
}

type ListGame struct {
	Id           primitive.ObjectID `json:"id" example:"507f1f77bcf86cd799439011"`
	Status       string             `json:"status" example:"waiting"`
	OwnerId      primitive.ObjectID `json:"owner_id" example:"507f1f77bcf86cd799439011"`
	Name         string             `json:"name" example:"test"`
	CountPlayers int                `json:"count_players" example:"3"`
	MaxPlayers   int                `json:"max_players" example:"4"`
	WithPassword bool               `json:"with_password" example:"true"`
}
