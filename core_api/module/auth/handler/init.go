package handler

import (
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/auth/service"
	sharedService "github.com/Brotiger/per-painted_poker-backend/core_api/shared/service"
)

type AuthHandler struct {
	AuthService         *service.AuthService
	RefreshTokenService *service.RefreshTokenService
	SharedTokenService  *sharedService.TokenService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		AuthService:         service.NewAuthService(),
		RefreshTokenService: service.NewRefreshTokenService(),
		SharedTokenService:  sharedService.NewTokenService(),
	}
}