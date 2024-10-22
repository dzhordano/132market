package config

import (
	"net"
	"os"

	"github.com/joho/godotenv"
)

const (
	httpHost = "HTTP_HOST"
	httpPort = "HTTP_PORT"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

type HttpConfig interface {
	Address() string
}

type httpConfig struct {
	address string
}

func NewHttpConfig() HttpConfig {
	host := os.Getenv(httpHost)
	if host == "" {
		panic("HTTP_HOST is not set")
	}

	port := os.Getenv(httpPort)
	if port == "" {
		panic("HTTP_PORT is not set")
	}

	return &httpConfig{
		address: net.JoinHostPort(host, port),
	}
}

func (c *httpConfig) Address() string {
	return c.address
}
