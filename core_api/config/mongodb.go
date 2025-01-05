package config

type MongoDB struct {
	Uri              string `env:"CORE_API_MONGODB_URI" envDefault:"mongodb://127.0.0.1:27017/"`
	Username         string `env:"CORE_API_MONGODB_USERNAME"`
	Password         string `env:"CORE_API_MONGODB_PASSWORD"`
	Database         string `env:"CORE_API_MONGODB_DATABASE" envDefault:"poker"`
	ConnectTimeoutMs int    `env:"CORE_API_MONGODB_CONNECT_TIMEOUT_MS" envDefault:"30000"`

	Table struct {
		User         string `env:"CORE_API_MONGODB_TABLE_USER" envDefault:"user"`
		Game         string `env:"CORE_API_MONGODB_TABLE_GAME" envDefault:"game"`
		Code         string `env:"CORE_API_MONGODB_TABLE_CODE" envDefault:"code"`
		RefreshToken string `env:"CORE_API_MONGODB_TABLE_REFRESH_TOKEN" envDefault:"refresh_token"`
		JoinToken    string `env:"CORE_API_MONGODB_TABLE_JOIN_TOKEN" envDefault:"join_token"`
		Player       string `env:"CORE_API_MONGODB_TABLE_PLAYER" envDefault:"player"`
	}
}
