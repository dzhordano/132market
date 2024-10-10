package config

import (
	"errors"
	"os"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func (c *pgConfig) DSN() string {
	return c.dsn
}

func NewPGConfig() (PGConfig, error) {
	cfg := &pgConfig{
		dsn: os.Getenv(PGDSN),
	}

	if cfg.dsn == "" {
		return nil, errors.New("PG_DSN must be set")
	}

	return cfg, nil
}
