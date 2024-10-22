package grpc

import (
	"context"
	"dzhordano/132market/services/users/internal/application/interfaces"
	"dzhordano/132market/services/users/pkg/pb/user_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserController struct {
	userService interfaces.UserService
	user_v1.UnimplementedUserServiceV1Server
}

func NewUserController(userService interfaces.UserService) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) CreateUser(ctx context.Context, request *user_v1.CreateUserRequest) (*user_v1.CreateUserResponse, error) {
	return nil, nil
}

func (u *UserController) UpdateUser(ctx context.Context, request *user_v1.UpdateUserRequest) (*user_v1.UpdateUserResponse, error) {
	return nil, nil
}

func (u *UserController) DeleteUser(ctx context.Context, request *user_v1.DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (u *UserController) FindUserById(ctx context.Context, request *user_v1.FindUserByIdRequest) (*user_v1.FindUserByIdResponse, error) {
	return nil, nil
}

func (u *UserController) FindAllUsers(ctx context.Context, request *emptypb.Empty) (*user_v1.FindAllUsersResponse, error) {
	return nil, nil
}
