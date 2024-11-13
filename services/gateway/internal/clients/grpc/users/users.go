package users

import (
	"context"
	"dzhordano/132market/services/users/pkg/pb/user_v1"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:55001"
)

type UserClient struct {
	c user_v1.UserServiceV1Client
}

func NewUserClient(ctx context.Context) user_v1.UserServiceV1Client {
	cc, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to grpc server: %s. Error: %s", address, err)
	}

	log.Printf("Connected to grpc server: %s", address)

	return user_v1.NewUserServiceV1Client(cc)
}
