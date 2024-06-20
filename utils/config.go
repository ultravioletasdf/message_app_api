package utils

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	UseLocalDatabase  bool   `env:"USE_LOCAL_DATABASE,required"`
	LocalDatabaseName string `env:"LOCAL_DATABASE_NAME" envDefault:"./local.db"`

	TursoDatabaseName string `env:"TURSO_DATABASE_NAME" envDefault:"./turso.db"`
	TursoUrl          string `env:"TURSO_URL"`
	TursoToken        string `env:"TURSO_TOKEN"`

	ListenAddress string `env:"LISTEN_ADDRESS" envDefault:":3000"`
}

func GetConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
