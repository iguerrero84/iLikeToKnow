package config

import (
	"context"
	"github.com/caarlos0/env/v11"
	"os"
)

// LoadConfig is a helper for loading a configuration struct
// It simply calls the automatic env library to populate fields with struct tags
// If there are secrets to be loaded, this function should be extended
func LoadConfig[T any](ctx context.Context) (*T, error) {
	LoadDBConfig()
	config, err := env.ParseAs[T]()
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadDBConfig() PostgresConfigImpl {

	// _ = godotenv.Load() I am not using .env files

	cfg := PostgresConfigImpl{
		PostgresDBHost_:     "localhost",
		PostgresDBPort_:     "5432",
		PostgresDBUser_:     "postgres",
		PostgresDBName_:     "events",
		PostgresDBPassword_: "admin",
	}

	// Also set env variables globally for the running process
	os.Setenv("POSTGRES_DB_HOST", cfg.PostgresDBHost_)
	os.Setenv("POSTGRES_DB_PORT", cfg.PostgresDBPort_)
	os.Setenv("POSTGRES_DB_USER", cfg.PostgresDBUser_)
	os.Setenv("POSTGRES_DB_NAME", cfg.PostgresDBName_)
	os.Setenv("POSTGRES_DB_PASSWORD", cfg.PostgresDBPassword_)

	return cfg
}
