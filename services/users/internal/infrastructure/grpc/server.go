package grpc

import (
	"log"
	"net"

	desc "github.com/dzhordano/132market/services/users/pkg/pb/user_v1"

	"github.com/dzhordano/132market/services/users/internal/application/interfaces"
	ctrl "github.com/dzhordano/132market/services/users/internal/interfaces/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
}

func NewServer(svc interfaces.UserService) *GRPCServer {
	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials())) // TODO сюда над токены при деплое

	reflection.Register(server)

	desc.RegisterUserServiceV1Server(server, ctrl.NewUserController(svc))

	return &GRPCServer{
		server: server,
	}
}

func (s *GRPCServer) Run(address string) error {
	list, err := net.Listen("tcp", address)

	if err != nil {
		return err
	}

	log.Printf("Starting grpc server on address: %s", address)

	return s.server.Serve(list)
}
