package config

type MongoDB struct {
	Uri       string `env:"MONGODB_URI" envDefault:"mongodb://127.0.0.1:27017/"`
	Username  string `env:"MONGODB_USERNAME"`
	Password  string `env:"MONGODB_PASSWORD"`
	Database  string `env:"MONGODB_DATABASE" envDefault:"poker"`
	TimeoutMs int    `env:"MONGODB_TIMEOUT_MS" envDefault:"30000"`

	Table struct {
		User         string `env:"MONGODB_TABLE_USER" envDefault:"user"`
		RefreshToken string `env:"MONGODB_TABLE_REFRESH_TOKEN" envDefault:"refresh_token"`
		Game         string `env:"MONGODB_TABLE_GAME" envDefault:"game"`
	}
}
