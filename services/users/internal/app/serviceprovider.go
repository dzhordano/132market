package app

import (
	"context"
	"dzhordano/132market/services/users/internal/api/users"
	"dzhordano/132market/services/users/internal/config"
	"dzhordano/132market/services/users/internal/repository"
	usersrepo "dzhordano/132market/services/users/internal/repository/users"
	"dzhordano/132market/services/users/internal/service"
	usersservice "dzhordano/132market/services/users/internal/service/users"
	"dzhordano/132market/services/users/pkg/databases/pg"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient     pg.DB
	usersRepo    repository.UsersRepo
	usersService service.UsersService
	usersImpl    *users.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			panic(err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			panic(err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) pg.DB {
	if s.dbClient == nil {
		pl, err := pg.NewPGPool(ctx, s.PGConfig().DSN())
		if err != nil {
			panic(err)
		}

		s.dbClient = pl
	}

	return s.dbClient
}

func (s *serviceProvider) UsersRepo(ctx context.Context) repository.UsersRepo {
	if s.usersRepo == nil {
		s.usersRepo = usersrepo.New(s.DBClient(ctx))
	}

	return s.usersRepo
}

func (s *serviceProvider) UsersService(ctx context.Context) service.UsersService {
	if s.usersService == nil {
		s.usersService = usersservice.New(s.UsersRepo(ctx))
	}

	return s.usersService
}

func (s *serviceProvider) UsersImpl(ctx context.Context) *users.Implementation {
	if s.usersImpl == nil {
		s.usersImpl = users.New(s.UsersService(ctx))
	}

	return s.usersImpl
}
