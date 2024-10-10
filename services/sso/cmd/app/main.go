package main

import (
	"dzhordano/132market/services/sso/pkg/logger"
	"log/slog"
	"os"
)

func main() {
	log := logger.MustTintSlogLogger(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	log.Info("Starting SSO service...")
}
