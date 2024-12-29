package config

type App struct {
	Jwt struct {
		Secret               string `env:"CORE_API_APP_JWT_SECRET"`
		AccessTokenExpireAt  int    `env:"CORE_API_APP_JWT_ACCESS_TOKEN_EXPIRE_AT_M"`
		RefreshTokenExpireAt int    `env:"CORE_API_APP_JWT_REFRESH_TOKEN_EXPIRE_AT_M"`
	}
}
