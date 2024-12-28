package middleware

import "github.com/Brotiger/per-painted_poker-backend/app/shared/service"

type AuthMiddleware struct {
	tokenService *service.TokenService
}

func NewShared() *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: service.NewTokenService(),
	}
}
