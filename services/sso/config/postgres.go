package config

import "os"

const (
	postgresDSN = "POSTGRES_DSN"
)

type PgConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func MustNewPostgresConfig() PgConfig {
	dsn := os.Getenv(postgresDSN)
	if dsn == "" {
		panic("POSTGRES_DSN is not set")
	}

	return &pgConfig{dsn: dsn}
}

func (c *pgConfig) DSN() string {
	return c.dsn
}
