package model

import "github.com/golang-jwt/jwt"

type ConnectClaims struct {
	GameId string `json:"gameId"`
	jwt.StandardClaims
}
