package main

import (
	"dzhordano/132market/services/users/internal/application/services"
	"dzhordano/132market/services/users/internal/infrastructure/db/postgres"
	"dzhordano/132market/services/users/internal/infrastructure/grpc"
	"dzhordano/132market/services/users/pkg/logger"
	"log/slog"
	"os"
)

func main() {
	logger := logger.NewTintSlogLogger(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // TODO не над хардкодить левел...
	})

	pool := postgres.NewPool("postgres://postgres:postgres@localhost:5401/postgres?sslmode=disable")

	repo := postgres.NewUserRepository(pool)

	svc := services.NewUserService(logger, repo)

	serv := grpc.NewServer(svc)

	logger.Info("Starting server ...")

	if err := serv.Run(":55001"); err != nil {
		slog.Error("Failed to run server:", slog.String("error", err.Error()))
	}

}
