package users

import (
	"context"
	"fmt"
	"log"

	user_v1 "github.com/dzhordano/132market/services/users/pkg/pb/user_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UsersClientInterface interface {
	CreateUser(ctx context.Context, req *user_v1.CreateUserRequest) (*user_v1.CreateUserResponse, error)
	FindUserByEmail(ctx context.Context, req *user_v1.FindUserByEmailRequest) (*user_v1.FindUserByEmailResponse, error)
	FindUserById(ctx context.Context, req *user_v1.FindUserByIdRequest) (*user_v1.FindUserByIdResponse, error)
}

type UsersClient struct {
	client user_v1.UserServiceV1Client
}

// FIXME ADD LOGGING INTERCEPTOR + RETRY INTERCEPTOR + TIMEOUT INTERCEPTOR ETC.
func NewUsersClient(addr string) (*UsersClient, error) {
	cc, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	log.Printf("connected to %s", addr)

	return &UsersClient{
		client: user_v1.NewUserServiceV1Client(cc),
	}, nil
}

func (c *UsersClient) CreateUser(ctx context.Context, req *user_v1.CreateUserRequest) (*user_v1.CreateUserResponse, error) {
	return c.client.CreateUser(ctx, req)
}

func (c *UsersClient) FindUserByEmail(ctx context.Context, req *user_v1.FindUserByEmailRequest) (*user_v1.FindUserByEmailResponse, error) {
	return c.client.FindUserByEmail(ctx, req)
}

func (c *UsersClient) FindUserById(ctx context.Context, req *user_v1.FindUserByIdRequest) (*user_v1.FindUserByIdResponse, error) {
	return c.client.FindUserById(ctx, req)
}
