package app

import (
	"context"
	"dzhordano/132market/services/users/internal/config"
	"dzhordano/132market/services/users/pkg/users_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.init(ctx)

	return a, err
}

func (a *App) init(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Run(_ context.Context) error {
	return a.runGRPCServer()
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	// TODO Заменить с использованием TLS.
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	users_v1.RegisterUsersV1Server(a.grpcServer, a.serviceProvider.UsersImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Println("Starting gRPC server on:", a.serviceProvider.GRPCConfig().Address())

	lis, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	return a.grpcServer.Serve(lis)
}
