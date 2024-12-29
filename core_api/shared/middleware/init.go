package middleware

import "github.com/Brotiger/poker-core_api/core_api/shared/service"

type AuthMiddleware struct {
	tokenService *service.TokenService
}

func NewShared() *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: service.NewTokenService(),
	}
}
