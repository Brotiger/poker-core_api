package model

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
