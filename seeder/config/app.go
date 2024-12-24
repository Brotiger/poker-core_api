package config

type App struct {
	Root struct {
		Username string `env:"APP_ROOT_USERNAME"`
		Email    string `env:"APP_ROOT_EMAIL"`
		Password string `env:"APP_ROOT_PASSWORD"`
	}
}
