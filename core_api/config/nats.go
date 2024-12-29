package config

type Nats struct {
	Addr                 string `env:"CORE_API_NATS_ADDR" envDefault:"localhost:4222"`
	ClientCert           string `env:"CORE_API_NATS_CLIENT_CERT"`
	ClientKey            string `env:"CORE_API_NATS_CLIENT_KEY"`
	RootCA               string `env:"CORE_API_NATS_ROOT_CA"`
	ReconnectWait        int    `env:"CORE_API_NATS_RECONNECT_WAIT" envDefault:"10000"`
	PingInterval         int    `env:"CORE_API_NATS_PING_INTERVAL" envDefault:"20000"`
	MaxReconnects        int    `env:"CORE_API_NATS_MAX_RECONNECTS" envDefault:"10"`
	RetryOnFailedConnect bool   `env:"CORE_API_NATS_RETRY_ON_FAILED_CONNECT" envDefault:"true"`
	MaxPingOutstanding   int    `env:"CORE_API_NATS_MAX_PING_OUTSTANDING" envDefault:"5"`
	Streams              struct {
		Mailer string `env:"CORE_API_NATS_STREAMS_MAILER" envDefault:"mailer"`
	}
}
