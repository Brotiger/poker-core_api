package config

type App struct {
	Jwt struct {
		Secret   string `env:"APP_JWT_SECRET"`
		TimeoutH int    `env:"APP_JWT_TIMEOUT_H"`
	}
}
