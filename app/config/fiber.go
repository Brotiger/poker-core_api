package config

type Fiber struct {
	Listen           string `env:"FIBER_LISTEN" envDefault:":8080"`
	RequestTimeoutMs int    `env:"FIBER_REQUEST_TIMEOUT_MS" envDefault:"3000"`
	Swagger          struct {
		Host string `env:"FIBER_SWAGGER_HOST" envDefault:"localhost:8080"`
	}
}
