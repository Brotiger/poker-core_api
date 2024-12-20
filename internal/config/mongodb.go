package config

type MongoDB struct {
	Uri       string `env:"MONGODB_URI" envDefault:"mongodb://127.0.0.1:27017/"`
	Username  string `env:"MONGODB_USERNAME"`
	Password  string `env:"MONGODB_PASSWORD"`
	Database  string `env:"MONGODB_DATABASE" envDefault:"poker"`
	TimeoutMs int    `env:"MONGODB_TIMEOUT_MS" envDefault:"30000"`
}
