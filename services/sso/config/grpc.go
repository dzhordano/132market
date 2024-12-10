package config

import (
	"net"
	"os"
)

// GRPC .env variables
const (
	grpcSSOHost = "GRPC_SSO_HOST"
	grpcSSOPort = "GRPC_SSO_PORT"

	grpcTokenHost = "GRPC_TOKEN_HOST"
	grpcTokenPort = "GRPC_TOKEN_PORT"

	grpcUsersHost = "GRPC_USERS_HOST"
	grpcUsersPort = "GRPC_USERS_PORT"
)

type GrpcConfig interface {
	Address() string
}

type grpcSsoConfig struct {
	host string
	port string
}

func MustNewGrpcSsoConfig() GrpcConfig {
	host := os.Getenv(grpcSSOHost)
	if host == "" {
		panic("GRPC_SSO_HOST is not set")
	}

	port := os.Getenv(grpcSSOPort)
	if port == "" {
		panic("GRPC_SSO_PORT is not set")
	}

	return &grpcSsoConfig{
		host: host,
		port: port,
	}
}

func (c *grpcSsoConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}

type grpcTokenConfig struct {
	host string
	port string
}

func MustNewGrpcTokenConfig() GrpcConfig {
	host := os.Getenv(grpcTokenHost)
	if host == "" {
		panic("GRPC_TOKEN_HOST is not set")
	}

	port := os.Getenv(grpcTokenPort)
	if port == "" {
		panic("GRPC_TOKEN_PORT is not set")
	}

	return &grpcTokenConfig{
		host: host,
		port: port,
	}
}

func (c *grpcTokenConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}

type grpcUsersConfig struct {
	host string
	port string
}

func MustNewGrpcUsersConfig() GrpcConfig {
	host := os.Getenv(grpcUsersHost)
	if host == "" {
		panic("GRPC_USERS_HOST is not set")
	}

	port := os.Getenv(grpcUsersPort)
	if port == "" {
		panic("GRPC_USERS_PORT is not set")
	}

	return &grpcUsersConfig{
		host: host,
		port: port,
	}
}

func (c *grpcUsersConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
