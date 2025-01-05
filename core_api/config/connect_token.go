package config

type ConnectToken struct {
	Length int `env:"CORE_API_CONNECT_TOKEN_LENGTH" envDefault:"32"`
}
