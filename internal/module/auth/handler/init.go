package handler

import "github.com/Brotiger/per-painted_poker-backend/internal/module/auth/service"

type Auth struct {
	AuthService         *service.Auth
	RefreshTokenService *service.RefreshToken
}

func NewAuth() *Auth {
	return &Auth{
		AuthService:         service.NewAuth(),
		RefreshTokenService: service.NewRefreshToken(),
	}
}
