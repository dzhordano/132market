package config

import "github.com/joho/godotenv"

// TODO Заменить на что-то красивее и по-практичнее?
var (
	loadPaths = []string{".env"}
)

const (
	GRPCHost = "GRPC_HOST" // GRPC_HOST is the host of the gRPC server
	GRPCPort = "GRPC_PORT" // GRPC_PORT is the port of the gRPC server

	PGDSN = "PG_DSN" // PG_DSN is the DSN of the PostgreSQL database that encapsulates connection parameters
)

func Load() error {
	if err := godotenv.Load(loadPaths...); err != nil {
		return err
	}

	return nil
}
