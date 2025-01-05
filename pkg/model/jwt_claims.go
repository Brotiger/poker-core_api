package model

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTClaims struct {
	UserId primitive.ObjectID `json:"userId"`
	jwt.StandardClaims
}
