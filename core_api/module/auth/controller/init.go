package controller

import (
	"github.com/Brotiger/poker-core_api/core_api/module/auth/service"
	sharedService "github.com/Brotiger/poker-core_api/pkg/service"
)

type AuthController struct {
	AuthService         *service.AuthService
	RefreshTokenService *service.RefreshTokenService
	SharedTokenService  *sharedService.TokenService
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthService:         service.NewAuthService(),
		RefreshTokenService: service.NewRefreshTokenService(),
		SharedTokenService:  sharedService.NewTokenService(),
	}
}
