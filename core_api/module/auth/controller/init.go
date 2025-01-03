package controller

import (
	"github.com/Brotiger/poker-core_api/core_api/module/auth/service"
	pkgService "github.com/Brotiger/poker-core_api/pkg/service"
)

type AuthController struct {
	AuthService         *service.AuthService
	RefreshTokenService *service.RefreshTokenService
	TokenService        *pkgService.TokenService
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthService:         service.NewAuthService(),
		RefreshTokenService: service.NewRefreshTokenService(),
		TokenService:        pkgService.NewTokenService(),
	}
}
