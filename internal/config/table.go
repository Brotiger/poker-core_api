package config

type Table struct {
	User string `env:"MONGODB_TABLE_USER" envDefault:"users"`
}
