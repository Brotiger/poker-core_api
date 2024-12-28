package model

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
