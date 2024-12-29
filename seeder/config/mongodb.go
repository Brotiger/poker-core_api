package config

type MongoDB struct {
	Uri              string `env:"SEEDER_MONGODB_URI" envDefault:"mongodb://127.0.0.1:27017/"`
	Username         string `env:"SEEDER_MONGODB_USERNAME"`
	Password         string `env:"SEEDER_MONGODB_PASSWORD"`
	Database         string `env:"SEEDER_MONGODB_DATABASE" envDefault:"poker"`
	ConnectTimeoutMs int    `env:"SEEDER_MONGODB_CONNECT_TIMEOUT_MS" envDefault:"30000"`
	QueryTimeoutMs   int    `env:"SEEDER_MONGODB_QUERY_TIMEOUT_MS" envDefault:"300"`

	Table struct {
		User string `env:"SEEDER_MONGODB_TABLE_USER" envDefault:"user"`
	}
}
