package config

import (
	"net"
	"os"
)

// GRPC .env variables
const (
	grpcHost = "GRPC_HOST"
	grpcPort = "GRPC_PORT"
)

type GrpcConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func MustNewGrpcConfig() GrpcConfig {
	host := os.Getenv(grpcHost)
	if host == "" {
		panic("GRPC_HOST is not set")
	}

	port := os.Getenv(grpcPort)
	if port == "" {
		panic("GRPC_PORT is not set")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}
}

func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
