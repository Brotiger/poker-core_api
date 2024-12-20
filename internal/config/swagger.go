package config

type Swagger struct {
	Host string `env:"SWAGGER_HOST" envDefault:"localhost:8080"`
}
