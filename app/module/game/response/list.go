package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type List struct {
	Total int64  `json:"total" example:"100"`
	Games []Game `json:"games"`
}

type Game struct {
	Id           primitive.ObjectID `json:"id" example:"507f1f77bcf86cd799439011"`
	Status       string             `json:"status" example:"waiting"`
	OwnerId      primitive.ObjectID `json:"ownerId" example:"507f1f77bcf86cd799439011"`
	Name         string             `json:"name" example:"test"`
	CountPlayers int                `json:"countPlayers" example:"3"`
	MaxPlayers   int                `json:"maxPlayers" example:"4"`
	WithPassword bool               `json:"withPassword" example:"true"`
}
