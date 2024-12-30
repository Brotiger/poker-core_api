package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConfirmedEmail struct {
	UserId primitive.ObjectID `json:"user_id"`
	Code   string             `json:"code"`
}
