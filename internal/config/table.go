package config

type Table struct {
	User         string `env:"TABLE_USER" envDefault:"user"`
	RefreshToken string `env:"TABLE_REFRESH_TOKEN" envDefault:"refresh_token"`
}
