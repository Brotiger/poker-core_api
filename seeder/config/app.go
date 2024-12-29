package config

type App struct {
	Root struct {
		Username string `env:"SEEDER_APP_ROOT_USERNAME"`
		Email    string `env:"SEEDER_APP_ROOT_EMAIL"`
		Password string `env:"SEEDER_APP_ROOT_PASSWORD"`
	}
}
