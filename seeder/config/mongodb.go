package config

type MongoDB struct {
	Uri              string `env:"MONGODB_URI" envDefault:"mongodb://127.0.0.1:27017/"`
	Username         string `env:"MONGODB_USERNAME"`
	Password         string `env:"MONGODB_PASSWORD"`
	Database         string `env:"MONGODB_DATABASE" envDefault:"poker"`
	ConnectTimeoutMs int    `env:"MONGODB_CONNECT_TIMEOUT_MS" envDefault:"30000"`
	QueryTimeoutMs   int    `env:"MONGODB_QUERY_TIMEOUT_MS" envDefault:"300"`

	Table struct {
		User string `env:"MONGODB_TABLE_USER" envDefault:"user"`
	}
}
