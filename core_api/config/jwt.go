package config

type JWT struct {
	Secret               string `env:"CORE_API_JWT_SECRET"`
	AccessTokenExpireAt  int    `env:"CORE_API_JWT_ACCESS_TOKEN_EXPIRE_AT_M"`
	RefreshTokenExpireAt int    `env:"CORE_API_JWT_REFRESH_TOKEN_EXPIRE_AT_M"`
}
