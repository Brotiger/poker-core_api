package config

type App struct {
	Jwt struct {
		Secret               string `env:"APP_JWT_SECRET"`
		AccessTokenExpireAt  int    `env:"APP_JWT_ACCESS_TOKEN_EXPIRE_AT_M"`
		RefreshTokenExpireAt int    `env:"APP_JWT_REFRESH_TOKEN_EXPIRE_AT_M"`
	}
}
