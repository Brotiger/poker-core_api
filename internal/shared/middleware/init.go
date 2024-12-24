package middleware

import "github.com/Brotiger/per-painted_poker-backend/internal/shared/service"

type Auth struct {
	ServiceToken *service.Token
}

func NewShared() *Auth {
	return &Auth{
		ServiceToken: service.NewToken(),
	}
}
