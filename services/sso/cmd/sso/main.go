package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/dzhordano/132market/services/sso/config"
	"github.com/dzhordano/132market/services/sso/internal/application/services"
	"github.com/dzhordano/132market/services/sso/internal/infrastructure/db/postgres"
	"github.com/dzhordano/132market/services/sso/internal/infrastructure/grpc"
	"github.com/dzhordano/132market/services/sso/pkg/hasher"
	"github.com/dzhordano/132market/services/sso/pkg/jwt"
	"github.com/dzhordano/132market/services/sso/pkg/logger"
)

func main() {
	fmt.Println("Hello, World!")

	// LOGGER

	log := logger.NewTintSlogLogger(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// CFG

	// DB

	db := postgres.NewPool(config.MustNewPostgresConfig().DSN())

	// REPO

	uRepo := postgres.NewUsersRepository(db)
	rRepo := postgres.NewRolesRepository(db)

	// PKG DEPS

	h := hasher.NewArgon2Hasher(
		config.MustNewArgon2Config().Time(),
		config.MustNewArgon2Config().SaltLen(),
		config.MustNewArgon2Config().Memory(),
		config.MustNewArgon2Config().Threads(),
		config.MustNewArgon2Config().KeyLen(),
	)

	tokenGenerator := jwt.NewJwtGenerator(
		config.MustNewJwtConfig().SigningKey(),
		config.MustNewJwtConfig().ATTL(),
		config.MustNewJwtConfig().RTTL(),
	)

	tokenValidator := jwt.NewJwtValidator(
		config.MustNewJwtConfig().SigningKey(),
		config.MustNewJwtConfig().ATTL(),
		config.MustNewJwtConfig().RTTL(),
	)

	tokens := jwt.NewJwtService(tokenGenerator, tokenValidator)

	// SERVICES

	uSvc := services.NewUsersService(log, uRepo)
	athSvc := services.NewAuthenticationService(log, uSvc, tokens, h)
	atrSvc := services.NewAuthorizationService(log, uSvc, rRepo)
	vSvc := services.NewTokenValidationService(log, tokenValidator)

	// SERVER + CONTROLLERS

	grpcSSOServer := grpc.NewSSOServer(grpc.SSOServices{
		AuthenticationService: athSvc,
		AuthorizationService:  atrSvc,
	})

	grpcTokenServer := grpc.NewTokenServer(grpc.TokenServices{
		TokensService: vSvc,
	})

	// RUN

	svrWg := sync.WaitGroup{}

	svrWg.Add(1)
	go func() {
		defer svrWg.Done()
		if err := grpcSSOServer.Run(config.MustNewGrpcSsoConfig().Address()); err != nil {
			log.Error("Failed to run grpc server:", slog.String("error", err.Error()))
		}
	}()

	svrWg.Add(1)
	go func() {
		defer svrWg.Done()
		if err := grpcTokenServer.Run(config.MustNewGrpcTokenConfig().Address()); err != nil {
			log.Error("Failed to run grpc server:", slog.String("error", err.Error()))
		}
	}()

	svrWg.Wait()

	// TODO GRACEFUL SHUTDOWN
}
