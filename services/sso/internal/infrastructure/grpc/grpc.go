package grpc

import (
	"log"
	"net"

	"github.com/dzhordano/132market/services/sso/internal/application/interfaces"
	"github.com/dzhordano/132market/services/sso/pkg/pb/sso_v1"
	"github.com/dzhordano/132market/services/sso/pkg/pb/validation_v1"

	ctrl "github.com/dzhordano/132market/services/sso/internal/interfaces/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type GRPCSSOServer struct {
	server *grpc.Server
}

type SSOServices struct {
	AuthenticationService interfaces.AuthenticationService
	AuthorizationService  interfaces.AuthorizationService
}

func NewSSOServer(svcs SSOServices) *GRPCSSOServer {
	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials())) // TODO сюда над токены при деплое

	sso_v1.RegisterAuthenticationV1Server(server, ctrl.NewAuthenticationController(
		svcs.AuthenticationService,
	))
	sso_v1.RegisterAuthorizationV1Server(server, ctrl.NewAuthorizationController(
		svcs.AuthorizationService,
	))

	reflection.Register(server)

	return &GRPCSSOServer{
		server: server,
	}
}

func (s *GRPCSSOServer) Run(address string) error {
	list, err := net.Listen("tcp", address)

	if err != nil {
		return err
	}

	log.Printf("Starting grpc server on address: %s", address)

	return s.server.Serve(list)
}

// GRPCTokenServer implements JWT token validation.
// It's separate due to its usage in different microservices, and therefore has a different address for convenience.
type GRPCTokenServer struct {
	server *grpc.Server
}

type TokenServices struct {
	TokensService interfaces.TokenValidationService
}

func NewTokenServer(svcs TokenServices) *GRPCTokenServer {
	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials())) // TODO сюда над токены при деплое. а надо ли..?

	validation_v1.RegisterValidationV1Server(server, ctrl.NewValidationController(
		svcs.TokensService,
	))

	reflection.Register(server)

	return &GRPCTokenServer{
		server: server,
	}
}

func (s *GRPCTokenServer) Run(address string) error {
	list, err := net.Listen("tcp", address)

	if err != nil {
		return err
	}

	log.Printf("Starting grpc server on address: %s", address)

	return s.server.Serve(list)
}
