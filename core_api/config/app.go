package config

type App struct {
	GracefulShutdownTimeoutMS int `env:"CORE_API_APP_GRACEFUL_SHUTDOWN_TIMEOUT" envDefault:"10000"`
	CodeLength                int `env:"CORE_API_APP_CODE_LENGTH" envDefault:"15"`
}
