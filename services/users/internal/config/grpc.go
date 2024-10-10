package config

import (
	"errors"
	"net"
	"os"
)

type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}

func NewGRPCConfig() (GRPCConfig, error) {
	cfg := &grpcConfig{
		host: os.Getenv(GRPCHost),
		port: os.Getenv(GRPCPort),
	}

	if cfg.host == "" {
		return nil, errors.New("GRPC_HOST must be set")
	}

	if cfg.port == "" {
		return nil, errors.New("GRPC_PORT must be set")
	}

	return cfg, nil
}
