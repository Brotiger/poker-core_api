package middleware

import "github.com/Brotiger/poker-core_api/pkg/service"

type AuthMiddleware struct {
	tokenService *service.TokenService
}

func NewShared() *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: service.NewTokenService(),
	}
}
