package grpc

import (
	"context"
	"dzhordano/132market/services/users/internal/application/interfaces"
	"dzhordano/132market/services/users/internal/interfaces/grpc/dto/mapper"
	"dzhordano/132market/services/users/pkg/pb/user_v1"

	"github.com/google/uuid"
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
	userCommand := mapper.CreateUserRequestToCommand(request)

	commandResp, err := u.userService.CreateUser(ctx, userCommand)
	if err != nil {
		return nil, err
	}

	response := mapper.ToUserResponse(commandResp.Result)

	return &user_v1.CreateUserResponse{User: response}, nil
}

func (u *UserController) UpdateUser(ctx context.Context, request *user_v1.UpdateUserRequest) (*user_v1.UpdateUserResponse, error) {
	userCommand := mapper.UpdateUserRequestToCommand(request)

	commandResp, err := u.userService.UpdateUser(ctx, userCommand)
	if err != nil {
		return nil, err
	}

	response := mapper.ToUserResponse(commandResp.Result)

	return &user_v1.UpdateUserResponse{User: response}, nil
}

func (u *UserController) DeleteUser(ctx context.Context, request *user_v1.DeleteUserRequest) (*emptypb.Empty, error) {
	id, err := uuid.FromBytes([]byte(request.Id))
	if err != nil {
		return nil, err
	}

	err = u.userService.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (u *UserController) FindUserById(ctx context.Context, request *user_v1.FindUserByIdRequest) (*user_v1.FindUserByIdResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	queryResp, err := u.userService.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	response := mapper.ToUserResponse(queryResp.Result)

	return &user_v1.FindUserByIdResponse{User: response}, nil
}

func (u *UserController) FindAllUsers(ctx context.Context, request *user_v1.FindAllUsersRequest) (*user_v1.FindAllUsersResponse, error) {
	queryResp, err := u.userService.FindAllUsers(ctx, request.GetLimit(), request.GetOffset())
	if err != nil {
		return nil, err
	}

	response := mapper.ToUserListResponse(queryResp.Result)

	return &user_v1.FindAllUsersResponse{Users: response}, nil
}