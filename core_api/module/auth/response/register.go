package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type Register struct {
	Id primitive.ObjectID `json:"id" example:"507f191e810c19729de860ea"`
}
