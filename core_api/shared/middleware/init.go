package middleware

import "github.com/Brotiger/per-painted_poker-backend/core_api/shared/service"

type AuthMiddleware struct {
	tokenService *service.TokenService
}

func NewShared() *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: service.NewTokenService(),
	}
}
